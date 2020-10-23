package command

import (
	"fmt"
	"time"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor/deployment"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/factory"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/urfave/cli"
)

const artifactTimeStampFormat = "20060102T150405Z"

type DeploymentBackupCommand struct {
}

func NewDeploymentBackupCommand() DeploymentBackupCommand {
	return DeploymentBackupCommand{}
}

func (d DeploymentBackupCommand) Cli() cli.Command {
	return cli.Command{
		Name:    "backup",
		Aliases: []string{"b"},
		Usage:   "Backup a deployment",
		Action:  d.Action,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "with-manifest",
				Usage: "Download the deployment manifest",
			},
			cli.StringFlag{
				Name:  "artifact-path, a",
				Usage: "Specify an optional path to save the backup artifacts to",
			},
			cli.BoolFlag{
				Name:  "unsafe-lock-free-backup",
				Usage: "Experimental feature to skip locking steps when backing up TAS. Use only on TAS and at your own risk",
			},
		},
	}
}

func (d DeploymentBackupCommand) Action(c *cli.Context) error {
	trapSigint(true)

	username, password, target, caCert, bbrVersion, debug, deployment, allDeployments := getDeploymentParams(c)
	withManifest := c.Bool("with-manifest")
	artifactPath := c.String("artifact-path")
	lockFree := c.Bool("unsafe-lock-free-backup")

	if allDeployments {
		return backupAll(target, username, password, caCert, artifactPath, withManifest, bbrVersion, debug, lockFree)
	}

	return backupSingleDeployment(deployment, target, username, password, caCert, artifactPath, withManifest, bbrVersion, debug, lockFree)
}

func backupAll(target, username, password, caCert, artifactPath string, withManifest bool, bbrVersion string, debug bool, lockFree bool) error {
	backupAction := func(deploymentName string) orchestrator.Error {
		timestamp := time.Now().UTC().Format(artifactTimeStampFormat)
		logFilePath, buffer, logger := createLogger(timestamp, artifactPath, deploymentName, debug)

		backuper, factoryErr := factory.BuildDeploymentBackuper(
			target,
			username,
			password,
			caCert,
			withManifest,
			bbrVersion,
			logger,
			timestamp,
			lockFree,
		)
		if factoryErr != nil {
			return orchestrator.NewError(factoryErr)
		}

		printlnWithTimestamp(fmt.Sprintf("Starting backup of %s, log file: %s", deploymentName, logFilePath))
		err := backuper.Backup(deploymentName, artifactPath)

		if err != nil {
			printlnWithTimestamp(fmt.Sprintf("ERROR: failed to backup %s", deploymentName))
			fmt.Println(buffer.String())
		} else {
			printlnWithTimestamp(fmt.Sprintf("Finished backup of %s", deploymentName))
		}

		return err
	}

	errorHandler := func(deploymentError deployment.AllDeploymentsError) error {
		if deployment.ContainsUnlockOrCleanup(deploymentError.DeploymentErrs) {
			return deploymentError.ProcessWithFooter(backupCleanupAllDeploymentsAdvisedNotice)
		}
		return deploymentError.Process()
	}

	fmt.Println("Starting backup...")

	logger, _ := factory.BuildBoshLoggerWithCustomBuffer(debug)
	boshClient, err := factory.BuildBoshClient(target, username, password, caCert, bbrVersion, logger)
	if err != nil {
		return processError(orchestrator.NewError(err))
	}

	return runForAllDeployments(backupAction,
		boshClient,
		"cannot be backed up",
		"backed up",
		errorHandler,
		deployment.NewParallelExecutor())
}

func backupSingleDeployment(deployment, target, username, password, caCert, artifactPath string, withManifest bool, bbrVersion string, debug bool, lockFree bool) error {
	logger := factory.BuildBoshLogger(debug)
	timeStamp := time.Now().UTC().Format(artifactTimeStampFormat)

	backuper, err := factory.BuildDeploymentBackuper(target, username, password, caCert, withManifest, bbrVersion, logger, timeStamp, lockFree)
	if err != nil {
		return processError(orchestrator.NewError(err))
	}

	backupErr := backuper.Backup(deployment, artifactPath)
	if backupErr.ContainsUnlockOrCleanupOrArtifactDirExists() {
		return processErrorWithFooter(backupErr, backupCleanupAdvisedNotice)
	}

	return processError(backupErr)
}

func printlnWithTimestamp(str string) {
	fmt.Printf("[%s] %s\n", time.Now().UTC().Format("15:04:05"), str)
}

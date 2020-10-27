package command

import (
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/cli/flags"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/factory"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/urfave/cli"
)

type DeploymentRestoreCommand struct {
}

func NewDeploymentRestoreCommand() DeploymentRestoreCommand {
	return DeploymentRestoreCommand{}
}

func (d DeploymentRestoreCommand) Cli() cli.Command {
	return cli.Command{
		Name:    "restore",
		Aliases: []string{"r"},
		Usage:   "Restore a deployment from backup",
		Action:  d.Action,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "artifact-path, a",
				Usage: "Path to the artifact to restore",
			},
			cli.BoolFlag{
				Name:  "unsafe-lock-free-restore",
				Usage: "Experimental feature to skip locking steps when restoring TAS. Use only on CF, with a backup that was also taken lock-free, and at your own risk",
			},
		},
	}
}

func (d DeploymentRestoreCommand) Action(c *cli.Context) error {
	trapSigint(false)

	if err := flags.Validate([]string{"artifact-path"}, c); err != nil {
		return err
	}

	deployment := c.Parent().String("deployment")
	artifactPath := c.String("artifact-path")
	lockFree := c.Bool("unsafe-lock-free-backup")

	restorer, err := factory.BuildDeploymentRestorer(c.Parent().String("target"),
		c.Parent().String("username"),
		c.Parent().String("password"),
		c.Parent().String("ca-cert"),
		c.App.Version,
		c.GlobalBool("debug"))

	if err != nil {
		return processError(orchestrator.NewError(err))
	}

	restoreErr := restorer.Restore(deployment, artifactPath, lockFree)
	return processError(restoreErr)
}

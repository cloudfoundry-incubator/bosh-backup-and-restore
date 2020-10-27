package orchestrator

import (
	"time"

	exe "github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"
)

func NewBackuper(backupManager BackupManager, logger Logger, deploymentManager DeploymentManager,
	lockOrderer LockOrderer, executor exe.Executor, nowFunc func() time.Time, artifactCopier ArtifactCopier, timestamp string, lockFree bool) *Backuper {

	findDeploymentStep := NewFindDeploymentStep(deploymentManager, logger)
	backupable := NewBackupableStep(lockOrderer, logger)
	createArtifact := NewCreateArtifactStep(logger, backupManager, deploymentManager, nowFunc, timestamp)
	lock := NewLockStep(lockOrderer, executor)

	backup := NewBackupStep(executor, lockFree)
	unlockAfterSuccessfulBackup := NewPostBackupUnlockStep(true, lockOrderer, executor)
	unlockAfterFailedBackup := NewPostBackupUnlockStep(false, lockOrderer, executor)
	drain := NewDrainStep(logger, artifactCopier)
	cleanup := NewCleanupStep()
	addFinishTimeStep := NewAddFinishTimeStep(nowFunc)

	workflow := NewWorkflow()
	workflow.StartWith(findDeploymentStep).OnSuccess(backupable)
	workflow.Add(backupable).OnSuccess(createArtifact).OnFailure(cleanup)
	workflow.Add(createArtifact).OnSuccess(lock).OnFailure(cleanup)
	workflow.Add(lock).OnSuccess(backup).OnFailure(unlockAfterFailedBackup)
	workflow.Add(backup).OnSuccess(unlockAfterSuccessfulBackup).OnFailure(unlockAfterFailedBackup)
	workflow.Add(unlockAfterSuccessfulBackup).OnSuccessOrFailure(drain)
	workflow.Add(unlockAfterFailedBackup).OnSuccessOrFailure(cleanup)
	workflow.Add(drain).OnSuccessOrFailure(cleanup)
	workflow.Add(cleanup).OnSuccessOrFailure(addFinishTimeStep)
	workflow.Add(addFinishTimeStep)

	return &Backuper{
		workflow: workflow,
	}
}

type Backuper struct {
	workflow *Workflow
}

type AuthInfo struct {
	Type   string
	UaaUrl string
}

//Backup checks if a deployment has backupable instances and backs them up.
func (b Backuper) Backup(deploymentName, artifactPath string) Error {
	session := NewSession(deploymentName)
	session.SetCurrentArtifactPath(artifactPath)

	err := b.workflow.Run(session)

	return err
}

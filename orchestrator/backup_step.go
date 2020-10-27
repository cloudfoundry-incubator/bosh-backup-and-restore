package orchestrator

import "github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"

type BackupStep struct {
	executor executor.Executor
	lockFree bool
}

func (s *BackupStep) Run(session *Session, _ bool) error {
	err := session.CurrentDeployment().Backup(s.executor, s.lockFree)
	if err != nil {
		return NewBackupError(err.Error())
	}
	return nil
}

func NewBackupStep(executor executor.Executor, lockFree bool) Step {
	return &BackupStep{
		executor: executor,
		lockFree: lockFree,
	}
}

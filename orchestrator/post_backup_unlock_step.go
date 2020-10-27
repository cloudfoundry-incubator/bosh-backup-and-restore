package orchestrator

import "github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"

type PostBackupUnlockStep struct {
	afterSuccessfulBackup bool
	lockOrderer           LockOrderer
	executor              executor.Executor
}

func NewPostBackupUnlockStep(afterSuccessfulBackup bool, lockOrderer LockOrderer, executor executor.Executor) Step {
	return &PostBackupUnlockStep{
		afterSuccessfulBackup: afterSuccessfulBackup,
		lockOrderer:           lockOrderer,
		executor:              executor,
	}
}

func (s *PostBackupUnlockStep) Run(session *Session, _ bool) error {
	err := session.CurrentDeployment().PostBackupUnlock(s.afterSuccessfulBackup, s.lockOrderer, s.executor)
	if err != nil {
		return NewPostUnlockError(err.Error())
	}
	return nil
}

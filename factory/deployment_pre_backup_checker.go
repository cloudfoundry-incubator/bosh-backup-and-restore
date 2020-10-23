package factory

import (
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/bosh"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orderer"
)

func BuildDeploymentBackupChecker(boshClient bosh.Client,
	logger bosh.Logger,
	withManifest bool) *orchestrator.BackupChecker {
	return orchestrator.NewBackupChecker(logger,
		bosh.NewDeploymentManager(boshClient, logger, withManifest), orderer.NewKahnBackupLockOrderer(false))
}

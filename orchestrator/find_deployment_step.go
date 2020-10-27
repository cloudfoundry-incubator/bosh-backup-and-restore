package orchestrator

type FindDeploymentStep struct {
	deploymentManager DeploymentManager
	logger            Logger
}

func NewFindDeploymentStep(deploymentManager DeploymentManager, logger Logger) Step {
	return &FindDeploymentStep{deploymentManager: deploymentManager, logger: logger}
}

func (s *FindDeploymentStep) Run(session *Session, lockFree bool) error {
	s.logger.Info("bbr", "Looking for scripts")
	deployment, err := s.deploymentManager.Find(session.DeploymentName(), lockFree)
	if err != nil {
		return err
	}

	session.SetCurrentDeployment(deployment)

	return nil
}

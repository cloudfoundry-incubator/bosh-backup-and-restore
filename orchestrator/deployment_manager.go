package orchestrator

//go:generate counterfeiter -o fakes/fake_deployment_manager.go . DeploymentManager
type DeploymentManager interface {
	Find(deploymentName string, lockFree bool) (Deployment, error)
	SaveManifest(deploymentName string, artifact Backup) error
}

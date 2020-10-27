package orchestrator

type CleanupPreviousStep struct{}

func NewCleanupPreviousStep() Step {
	return &CleanupPreviousStep{}
}

func (s *CleanupPreviousStep) Run(session *Session, _ bool) error {
	return session.CurrentDeployment().CleanupPrevious()
}

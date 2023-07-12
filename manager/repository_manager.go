package manager

// import "enigmacamp.com/final-project/team-4/track-prosto/repository"

type RepositoryManager interface {
	
	// CustomerRepo() repository.CustomerRepository
}

type repositoryManager struct {
	infra InfraManager
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return repositoryManager{
		infra: manager,
	}
}

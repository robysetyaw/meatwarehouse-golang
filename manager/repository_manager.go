package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager

	userRepo repository.UserRepository
}

var onceLoadUserRepo sync.Once
var onceLoadFarmerRepo sync.Once
var onceLoadPlantRepo sync.Once
var onceLoadBillRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.userRepo = repository.NewUserRepository(rm.infraManager.GetDB())
	})
	return rm.userRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}

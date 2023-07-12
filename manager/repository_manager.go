package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepository
	GetDailyExpenditureRepo() repository.DailyExpenditureRepository
}

type repoManager struct {
	infraManager InfraManager

	userRepo repository.UserRepository
	dailyExpenditureRepo repository.DailyExpenditureRepository
}

var onceLoadUserRepo sync.Once
var onceLoadDailyExpenditureRepo sync.Once
var onceLoadPlantRepo sync.Once
var onceLoadBillRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.userRepo = repository.NewUserRepository(rm.infraManager.GetDB())
	})
	return rm.userRepo
}

func (rm *repoManager) GetDailyExpenditureRepo() repository.DailyExpenditureRepository {
	onceLoadDailyExpenditureRepo.Do(func() {
		rm.dailyExpenditureRepo = repository.NewDailyExpenditureRepository(rm.infraManager.GetDB())
	})
	return rm.dailyExpenditureRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}

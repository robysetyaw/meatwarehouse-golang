package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/repository"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepository
	GetMeatRepo() repository.MeatRepository
	GetDailyExpenditureRepo() repository.DailyExpenditureRepository
	GetCustomerRepo() repository.CustomerRepository
	GetCompanyRepo() repository.CompanyRepository
}

type repoManager struct {
	infraManager         InfraManager
	customerRepo         repository.CustomerRepository
	userRepo             repository.UserRepository
	companyRepo          repository.CompanyRepository
	meatRepo             repository.MeatRepository
	dailyExpenditureRepo repository.DailyExpenditureRepository
}

var onceLoadUserRepo sync.Once
var onceLoadDailyExpenditureRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadCompanyRepo sync.Once
var onceLoadMeatRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.userRepo = repository.NewUserRepository(rm.infraManager.GetDB())
	})
	return rm.userRepo
}
func (rm *repoManager) GetCustomerRepo() repository.CustomerRepository {
	onceLoadCustomerRepo.Do(func() {
		rm.customerRepo = repository.NewCustomerRepository(rm.infraManager.GetDB())
	})
	return rm.customerRepo
}
func (rm *repoManager) GetCompanyRepo() repository.CompanyRepository {
	onceLoadCompanyRepo.Do(func() {
		rm.companyRepo = repository.NewCompanyRepository(rm.infraManager.GetDB())
	})
	return rm.companyRepo
}

func (rm *repoManager) GetMeatRepo() repository.MeatRepository {
	onceLoadMeatRepo.Do(func() {
		rm.meatRepo = repository.NewMeatRepository(rm.infraManager.GetDB())
	})
	return rm.meatRepo
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

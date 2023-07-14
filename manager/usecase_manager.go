package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
	GetLoginUsecase() usecase.LoginUseCase
	GetDailyExpenditureUsecase() usecase.DailyExpenditureUseCase
	GetCustomerUsecase() usecase.CustomerUseCase
	GetCompanyUsecase() usecase.CompanyUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	userUsecase             usecase.UserUseCase
	loginUsecase            usecase.LoginUseCase
	customerUsecase         usecase.CustomerUseCase
	companyUsecase          usecase.CompanyUseCase
	dailyExpenditureUsecase usecase.DailyExpenditureUseCase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadCompanyUsecase sync.Once
var onceLoadDailyExpenditureUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}
func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.loginUsecase = usecase.NewLoginUseCase(um.repoManager.GetUserRepo())
	})
	return um.loginUsecase
}
func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUseCase {
	onceLoadCustomerUsecase.Do(func() {
		um.customerUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo(), um.repoManager.GetCompanyRepo())
	})
	return um.customerUsecase
}
func (um *usecaseManager) GetCompanyUsecase() usecase.CompanyUseCase {
	onceLoadCompanyUsecase.Do(func() {
		um.companyUsecase = usecase.NewCompanyUseCase(um.repoManager.GetCompanyRepo())
	})
	return um.companyUsecase
}

func (um *usecaseManager) GetDailyExpenditureUsecase() usecase.DailyExpenditureUseCase {
	onceLoadDailyExpenditureUsecase.Do(func() {
		um.dailyExpenditureUsecase = usecase.NewDailyExpenditureUseCase(um.repoManager.GetDailyExpenditureRepo())
	})
	return um.dailyExpenditureUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

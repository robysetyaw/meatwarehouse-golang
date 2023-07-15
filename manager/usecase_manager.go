package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
	GetMeatUsecase() usecase.MeatUseCase
	GetLoginUsecase() usecase.LoginUseCase
	GetDailyExpenditureUsecase() usecase.DailyExpenditureUseCase
	GetReportUsecase() usecase.ReportUseCase
	GetCustomerUsecase() usecase.CustomerUseCase
	GetCompanyUsecase() usecase.CompanyUseCase
	GetTransactionUseCase() usecase.TransactionUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	userUsecase             usecase.UserUseCase
	meatUsecase             usecase.MeatUseCase
	loginUsecase            usecase.LoginUseCase
	customerUsecase         usecase.CustomerUseCase
	companyUsecase          usecase.CompanyUseCase
	dailyExpenditureUsecase usecase.DailyExpenditureUseCase
	reportUsecase           usecase.ReportUseCase
	transactionUseCase      usecase.TransactionUseCase
}

var onceLoadUserUsecase sync.Once
var onceLoadMeatUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadCompanyUsecase sync.Once
var onceLoadDailyExpenditureUsecase sync.Once
var onceLoadReportUsecase sync.Once
var onceLoadTxUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}

func (mm *usecaseManager) GetMeatUsecase() usecase.MeatUseCase {
	onceLoadMeatUsecase.Do(func() {
		mm.meatUsecase = usecase.NewMeatUseCase(mm.repoManager.GetMeatRepo())
	})
	return mm.meatUsecase
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
		um.dailyExpenditureUsecase = usecase.NewDailyExpenditureUseCase(um.repoManager.GetDailyExpenditureRepo(), um.repoManager.GetUserRepo())
	})
	return um.dailyExpenditureUsecase
}

func (um *usecaseManager) GetReportUsecase() usecase.ReportUseCase {
	onceLoadReportUsecase.Do(func() {
		um.reportUsecase = usecase.NewReportUseCase(um.repoManager.GetDailyExpenditureRepo())
	})
	return um.reportUsecase
}

func (um *usecaseManager) GetTransactionUseCase() usecase.TransactionUseCase {
	onceLoadTxUsecase.Do(func() {
		um.transactionUseCase = usecase.NewTransactionUseCase(
			um.repoManager.GetTransactionRepo(),
			um.repoManager.GetCustomerRepo(),
			um.repoManager.GetMeatRepo(),
			um.repoManager.GetCompanyRepo())
	})
	return um.transactionUseCase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

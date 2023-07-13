package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
	GetLoginUsecase() usecase.LoginUseCase
	GetDailyExpenditureUsecase() usecase.DailyExpenditureUseCase
	GetReportUsecase() usecase.ReportUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	userUsecase    usecase.UserUseCase
	loginUsecase	usecase.LoginUseCase
	dailyExpenditureUsecase usecase.DailyExpenditureUseCase
	reportUsecase usecase.ReportUseCase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadDailyExpenditureUsecase sync.Once
var onceLoadReportUsecase sync.Once

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

func (um *usecaseManager) GetDailyExpenditureUsecase() usecase.DailyExpenditureUseCase {
	onceLoadDailyExpenditureUsecase.Do(func() {
		um.dailyExpenditureUsecase = usecase.NewDailyExpenditureUseCase(um.repoManager.GetDailyExpenditureRepo())
	})
	return um.dailyExpenditureUsecase
}

func (um *usecaseManager) GetReportUsecase() usecase.ReportUseCase {
	onceLoadReportUsecase.Do(func() {
		um.reportUsecase = usecase.NewReportUseCase(um.repoManager.GetDailyExpenditureRepo())
	})
	return um.reportUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

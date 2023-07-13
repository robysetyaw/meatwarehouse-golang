package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
<<<<<<< HEAD
	GetLoginUsecase() usecase.LoginUsecase
}

type usecaseManager struct {
	repoManager  RepoManager
	userUsecase  usecase.UserUseCase
	loginUsecase usecase.LoginUsecase
}

var onceLoadLoginUsecase sync.Once
var onceLoadUserUsecase sync.Once

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUsecase {
	onceLoadLoginUsecase.Do(func() {
		um.loginUsecase = usecase.NewLoginUsecase(um.repoManager.GetUserRepo())
	})
	return um.loginUsecase
}
=======
	GetLoginUsecase() usecase.LoginUseCase
	GetDailyExpenditureUsecase() usecase.DailyExpenditureUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	userUsecase    usecase.UserUseCase
	loginUsecase	usecase.LoginUseCase
	dailyExpenditureUsecase usecase.DailyExpenditureUseCase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadDailyExpenditureUsecase sync.Once
>>>>>>> origin/dev-borr

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}
<<<<<<< HEAD
=======
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
>>>>>>> origin/dev-borr

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

package manager

import (
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/usecase"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
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

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

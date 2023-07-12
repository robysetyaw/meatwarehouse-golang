package delivery

import (
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/controller"
	"enigmacamp.com/final-project/team-4/track-prosto/manager"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	srv            *gin.Engine
}

func (s *server) Run() {
	s.srv.Use()
	controller.NewUserController(s.srv, s.usecaseManager.GetUserUsecase())
	controller.NewLoginController(s.srv, s.usecaseManager.GetLoginUsecase())
	controller.NewDailyExpenditureController(s.srv, s.usecaseManager.GetDailyExpenditureUsecase())
	s.srv.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
	}

}

package delivery

import (
<<<<<<< HEAD
	"enigmacamp.com/final-project/team-4/track-prosto/config"
=======
>>>>>>> origin/dev-borr
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/controller"
	"enigmacamp.com/final-project/team-4/track-prosto/manager"
	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UsecaseManager
	engine         *gin.Engine
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run()
	if err != nil {
		panic(err)
	}
}
func (s *Server) initController() {
	controller.NewUserController(s.engine, s.useCaseManager.GetUserUsecase())
	controller.NewLoginController(s.engine, s.useCaseManager.GetLoginUsecase())
}
func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)
	return &Server{useCaseManager: usecase, engine: r}
}

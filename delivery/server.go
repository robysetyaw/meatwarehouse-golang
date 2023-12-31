package delivery

import (
	"enigmacamp.com/final-project/team-4/track-prosto/config"
	"enigmacamp.com/final-project/team-4/track-prosto/delivery/controller"
	"enigmacamp.com/final-project/team-4/track-prosto/manager"
	"github.com/gin-contrib/cors"
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
	controller.NewMeatController(s.engine, s.useCaseManager.GetMeatUsecase())
	controller.NewLoginController(s.engine, s.useCaseManager.GetLoginUsecase())
	controller.NewCompanyController(s.engine, s.useCaseManager.GetCompanyUsecase())
	controller.NewCustomerController(s.engine, s.useCaseManager.GetCustomerUsecase())
	controller.NewDailyExpenditureController(s.engine, s.useCaseManager.GetDailyExpenditureUsecase())
	controller.NewReportController(s.engine, s.useCaseManager.GetReportUsecase())
	controller.NewTransactionController(s.engine, s.useCaseManager.GetTransactionUseCase())
	controller.NewCreditPaymentController(s.engine, s.useCaseManager.GetCreditPaymentUseCase())
	controller.NewStockMovementController(s.engine, s.useCaseManager.GetStockMovementReportUseCase())
}
func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	
	r := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(configCors))

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)
	return &Server{useCaseManager: usecase, engine: r}
}

package delivery

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func (s *Server) Run() {
	err := s.engine.Run()
	if err != nil {
		panic(err)
	}
}
func NewServer() *Server {
	r := gin.Default()
	return &Server{engine: r}
}

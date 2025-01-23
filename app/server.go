package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shyfeng955/go-clean/api"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *api.Router
}

func (s *Server) Start() {
	s.apiRouter.Routers(s.engine)
	if err := s.engine.Run(); err != nil {
		return
	}
}

func NewServer(engine *gin.Engine, apiRouter *api.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

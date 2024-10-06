package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/magavales/Music-Library/todo/internal/models/configs"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	config     configs.ServerConfig
}

func New() *Server {
	return &Server{httpServer: nil}
}

func (s *Server) Run(router *gin.Engine) error {
	err := s.config.Parse()
	if err != nil {
		return err
	}
	s.httpServer = &http.Server{
		Addr:    ":" + s.config.Get(),
		Handler: router,
	}
	return s.httpServer.ListenAndServe()
}

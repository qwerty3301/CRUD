package server

import (
	"MyCRUD/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	Port    string
	Handler handlers.UserHandler
}

func (s *Server) Create() *gin.Engine {
	router := gin.Default()
	s.registerRoutes(router)
	return router
}

func (s *Server) registerRoutes(router *gin.Engine) {
	auth := router.Group("/auth", handlers.UserIdentity)
	{
		auth.POST("/delete", s.Handler.DeleteHandler)
		auth.POST("/update", s.Handler.UpdateHandler)
	}
	router.POST("/authorization", s.Handler.Authorization)
	router.POST("/register", s.Handler.RegisterHandler)
}

func (s *Server) Start(router *gin.Engine) {
	err := router.Run(s.Port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

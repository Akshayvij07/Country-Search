package api

import (
	"github.com/Akshayvij07/country-search/internals/api/handler"
	"github.com/Akshayvij07/country-search/internals/api/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	Port   string
}

func (s *Server) Serve() error {
	return s.router.Run(":" + s.Port)
}

type api interface {
	Register(rg *gin.RouterGroup, appName string)
}

func NewServer(port string, handler *handler.Handler) *Server {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	basepath := router.Group("/api/countries")
	routes.Routes(basepath, handler)

	return &Server{
		router: router,
		Port:   port,
	}
}

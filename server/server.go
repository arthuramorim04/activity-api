package server

import (
	"log"

	routes "github.com/arthuramorim04/go-activity-api.git/server/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	router := gin.Default()
	router.Use(cors.Default())
	return Server{
		port:   "8080",
		server: router,
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoute(s.server)
	log.Print("Server run on port " + s.port + "...")
	log.Fatal(routes.Run(": " + s.port))
}

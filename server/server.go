package server

import (
	"log"

	routes "github.com/arthuramorim04/go-books-api.git/server/routers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoute(s.server)
	log.Print("Server run on port " + s.port + "...")
	log.Fatal(routes.Run(": " + s.port))
}

package server

import (
	"github.com/andrewscarlos/golang-book-api/server/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server is runnig on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}

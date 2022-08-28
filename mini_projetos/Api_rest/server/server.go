package server

import (
	"log"

	"github.com/JPaulo-Moura/16-api-rest/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

//Retorna a instancia de um servidor
func NewServer() Server {
	return Server{
		port:   "3333",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}

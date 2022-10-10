package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	// Add Routes
	router.POST("/participant", CreateParticipant)
	router.GET("/participants", GetParticipants)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
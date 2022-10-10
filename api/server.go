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
	router.GET("/participant/:refNumber", GetParticipant)
	router.PUT("/participant/:refNumber", UpdateParticipant)
	router.DELETE("/participant/:refNumber", DeleteParticipant)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
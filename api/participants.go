package api

import (
	"net/http"
	"participant-project/models"
	"participant-project/util"
	"time"

	"github.com/gin-gonic/gin"
)

var now = time.Now()
var participants = []models.Participant{}

func CreateParticipant(ctx *gin.Context) {
	var payload *models.CreateParticipantRequest

	// hnadle errors
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// create a new participant
	newParticipant := models.Participant{
		Id: util.GenerateUUID(),
		Name: payload.Name,
		DateOfBirth: payload.DateOfBirth,
		PhoneNumber: payload.PhoneNumber,
		Address: payload.Address,
		Ref: util.RandomReferenceNumber(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Append to the list since we are not persisting at this moment
	participants = append(participants, newParticipant)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Participant Created", "data": newParticipant})
}

func GetParticipants(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Participants retrieved", "results": len(participants), "data": participants})
}
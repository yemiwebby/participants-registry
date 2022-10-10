package api

import (
	"net/http"
	"participant-project/models"
	"participant-project/util"
	"time"

	"github.com/gin-gonic/gin"
)

var now = time.Now()
var participants = []models.Participant{
	{
		Id: "1", 
		Name: "John doe", 
		DateOfBirth: "20-29-20", 
		PhoneNumber: "84927742", 
		Address: "Texas", 
		Ref: "ED34", 
		CreatedAt: now, 
		UpdatedAt: now,
	},
}

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

// get a single participant
func GetParticipant(ctx *gin.Context) {
	participantRef := ctx.Param("refNumber")

	// Check if no params was included
	e, participant, _ := checkIfParticipantExists(participantRef)

	if !e {
		ctx.JSON(http.StatusNotFound, gin.H{"status":"fail", "message": "No participant with that reference number exists"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Participant retrieved", "data": participant})
	}
}

func UpdateParticipant(ctx *gin.Context) {
	participantRef := ctx.Param("refNumber")
	var payload *models.UpdateParticipant

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	e, updatedParticipant, p := checkIfParticipantExists(participantRef)
	
	if !e {
		ctx.JSON(http.StatusNotFound, gin.H{"status":"fail", "message": "No participant with that reference number exists"})
	} else {
		now := time.Now()
		updatedParticipant.Name = payload.Name
		updatedParticipant.Address = payload.Address
		updatedParticipant.DateOfBirth = payload.DateOfBirth
		updatedParticipant.PhoneNumber = payload.PhoneNumber
		updatedParticipant.UpdatedAt = now

		participants[p] = updatedParticipant

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Participant Updated", "data": updatedParticipant})
	}
}


func DeleteParticipant(ctx *gin.Context) {
	participantRef := ctx.Param("refNumber")

	e, _, p := checkIfParticipantExists(participantRef)

	if !e {
		ctx.JSON(http.StatusNotFound, gin.H{"status":"fail", "message": "No participant with that reference number exists"})
	} else {
		participants = append(participants[:p], participants[p+1:]...)
		ctx.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "Participant deleted", "data": participants})
	}

}

func checkIfParticipantExists(r string) (bool, models.Participant, int) {
	var exists bool = false
	var participant models.Participant
	var position int

	for i, a := range participants {
		if (a.Ref == r) {
			exists = true
			participant = a
			position = i
		}
	}

	return exists, participant, position
}
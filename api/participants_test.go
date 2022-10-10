package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"participant-project/models"
	"participant-project/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    return router
}

func TestCreateParticipant(t *testing.T) {
    r := SetUpRouter()
    // Setup route
    r.POST("/participant", CreateParticipant)

	newParticipant := models.Participant{
		Id: util.GenerateUUID(),
		Name: "Demo name",
		DateOfBirth: "2020-02-09",
		PhoneNumber: "000888444",
		Address: "Edinburgh Scotland",
		Ref: util.RandomReferenceNumber(),
		CreatedAt: now,
		UpdatedAt: now,
	}


    // Encode struct into JSON
    jsonVal, _ := json.Marshal(newParticipant)

    // Create the mock request you would like to test
    req, err := http.NewRequest(http.MethodPost, "/participant", bytes.NewBuffer(jsonVal))

    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)


    assert.Equal(t, http.StatusCreated, w.Code)

    if w.Code == http.StatusCreated {
        t.Logf("Expected to get status %d is the same as %d\n", http.StatusCreated, w.Code)
    }
    
}

func TestGetParticipants(t *testing.T) {
    r := SetUpRouter()
    r.GET("/participants", GetParticipants)

    req, err := http.NewRequest(http.MethodGet, "/participants", nil)

    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    participants := []models.Participant{
        {
            Id: util.GenerateUUID(), 
            Name: "Olususi Oluyemi", 
            DateOfBirth: "2020-02-09",
            PhoneNumber: "928824782", 
            Address: "Edinburgh Scotland", 
            Ref: util.RandomReferenceNumber(), 
            CreatedAt: now, 
            UpdatedAt: now,
        },
    }

    json.Unmarshal(w.Body.Bytes(), &participants)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, participants)
}
package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func TestHomepageHandler(t *testing.T) {
    mockResponse := `{"message":"Welcome to participant registry microservice"}`
    r := SetUpRouter()
    r.GET("/", HomepageHandler)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
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

func TestGetParticipant(t *testing.T) {
    r := SetUpRouter()
    r.GET("/participant/:refNumber", GetParticipant)

    req, err := http.NewRequest(http.MethodGet, "/participant/ED34", nil)

    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)


    pNotFound, _ := http.NewRequest(http.MethodPut, "/participant/KH32", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, pNotFound)
    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateParticipant(t *testing.T) {
    r := SetUpRouter()
    r.PUT("/participant/:refNumber", UpdateParticipant)

    // Update participant details
    updatedParticipant := models.Participant{
		Name: "New name",
		DateOfBirth: "2020-02-09",
		PhoneNumber: "000888444",
		Address: "London, England",
		UpdatedAt: now,
	}

    val, _ := json.Marshal(updatedParticipant)

    pFound, _ := http.NewRequest(http.MethodPut, "/participant/ED34", bytes.NewBuffer(val))
    w := httptest.NewRecorder()
    
    r.ServeHTTP(w, pFound)
    assert.Equal(t, http.StatusOK, w.Code)

    pNotFound, _ := http.NewRequest(http.MethodPut, "/participant/KH32", bytes.NewBuffer(val))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, pNotFound)
    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteParticipant(t *testing.T) {
    r := SetUpRouter()
    r.DELETE("/participant/:refNumber", DeleteParticipant)

    req, _ := http.NewRequest(http.MethodDelete, "/participant/ED34", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNoContent, w.Code)

    pNotFound, _ := http.NewRequest(http.MethodPut, "/participant/KH32", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, pNotFound)
    assert.Equal(t, http.StatusNotFound, w.Code)
}
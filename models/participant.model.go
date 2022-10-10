package models

import (
	"time"
)


type Participant struct {
	Id string `json:"id"`
	Name string `json:"name" binding:"required"`
	DateOfBirth string `json:"dob" binding:"required"`
	PhoneNumber	string `json:"phoneNumber" binding:"required"`
	Address	string `json:"address"`
	Ref string `json:"refNumber" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}


type CreateParticipantRequest struct {
	Name	string `json:"name" binding:"required"`
	DateOfBirth string `json:"dob" binding:"required"`
	PhoneNumber	string `json:"phoneNumber" binding:"required"`
	Address	string `json:"address"`
	Ref string `json:"refNumber"`
	CreatedAt	time.Time `json:"created_at,omitempty"`
	UpdatedAt	time.Time `json:"updated_at,omitempty"`
}

type UpdateParticipant struct {
	Name	string `json:"name" binding:"required"`
	DateOfBirth string `json:"dob" binding:"required"`
	PhoneNumber	string `json:"phoneNumber" binding:"required"`
	Address	string `json:"address"`
	Ref string `json:"refNumber"`
	CreatedAt	time.Time `json:"created_at,omitempty"`
	UpdatedAt	time.Time `json:"updated_at,omitempty"`
}

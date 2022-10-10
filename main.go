package main

import (
	"fmt"
	"log"
	"participant-project/api"
)


func main() {
	fmt.Println("Welcome to participants Registry Microservice")

	server := api.NewServer()

	err :=  server.Start("0.0.0.0:8080")

	if err != nil {
		log.Fatal("Cannot start the server:", err)
	}
}
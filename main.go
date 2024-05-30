package main

import (
	"log"

	"github.com/cqllum/db2ch/api"
)

func main() {
	// Initialize the router
	router := api.SetupRouter()

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

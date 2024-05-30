package main

import (
	"log"

	"github.com/cqllum/db2ch/api"
	"github.com/cqllum/db2ch/replication"
)

func main() {
	// Check MySQL connection
	if err := replication.CheckMySQLConnection("config/config.json"); err != nil {
		log.Printf("Failed to check MySQL connection: %v", err)
	}

	// Check MSSQL connection
	if err := replication.CheckMSSQLConnection("config/config.json"); err != nil {
		log.Printf("Failed to check MSSQL connection: %v", err)
	}

	// Initialize the router and start the server
	router := api.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

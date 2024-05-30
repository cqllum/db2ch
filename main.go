package main

import (
	"log"
	"net/http"

	"github.com/cqllum/db2ch/api"
)

func main() {
	// Your initialization code here...

	router := api.SetupRouter()

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

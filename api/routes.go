package api

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/start", StartHandler).Methods("POST")
	router.HandleFunc("/stop", StopHandler).Methods("POST")
	router.HandleFunc("/status", StatusHandler).Methods("GET")
	return router
}

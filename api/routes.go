package api

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/start", StartHandler).Methods("GET")
	router.HandleFunc("/stop", StopHandler).Methods("GET")
	router.HandleFunc("/status", StatusHandler).Methods("GET")
	return router
}

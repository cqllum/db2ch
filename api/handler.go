package api

import (
	"encoding/json"
	"net/http"

	"github.com/cqllum/db2ch/replication"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	replication.StartReplication()
	json.NewEncoder(w).Encode("Replication started")
}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	replication.StopReplication()
	json.NewEncoder(w).Encode("Replication stopped")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := replication.GetStatus()
	json.NewEncoder(w).Encode(status)
}

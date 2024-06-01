package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/cqllum/db2ch/config"
	"github.com/cqllum/db2ch/replication"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	dbType := r.URL.Query().Get("db_type")
	dbAlias := r.URL.Query().Get("db_alias")

	// Load the configuration
	config.LoadConfig()

	// Access the configuration based on dbType and dbAlias
	var dbConfig *config.DBConfig
	switch strings.ToLower(dbType) {
	case "mysql":
		dbConfig = findDBConfig(config.AppConfig.Databases.MySQL, dbAlias)
	case "postgresql":
		dbConfig = findDBConfig(config.AppConfig.Databases.PostgreSQL, dbAlias)
	case "mssql":
		dbConfig = findDBConfig(config.AppConfig.Databases.MSSQL, dbAlias)
	default:
		log.Printf("Unsupported database type: %s\n", dbType)
		http.Error(w, "Unsupported database type", http.StatusBadRequest)
		return
	}

	if dbConfig != nil {

		// Declare that the configuration was found
		log.Printf("Found %s configuration: %+v\n", dbType, *dbConfig)

		// Add your replication logic here using dbConfig
		err := replication.ReplicationInit(dbType, *dbConfig)
		if err != nil {
			log.Printf("Error in ReplicationService: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	} else {
		log.Printf("%s configuration with alias '%s' not found\n", dbType, dbAlias)
		http.Error(w, "Database configuration not found", http.StatusNotFound)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("StartHandler executed successfully"))
}

// findDBConfig finds a DBConfig with the given alias in the provided slice
func findDBConfig(dbConfigs []config.DBConfig, dbAlias string) *config.DBConfig {
	for _, dbConfig := range dbConfigs {
		if dbConfig.Name == dbAlias {
			return &dbConfig
		}
	}
	return nil
}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	replication.StopReplication()
	json.NewEncoder(w).Encode("Replication stopped")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := replication.GetStatus()
	json.NewEncoder(w).Encode(status)
}

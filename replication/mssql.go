package replication

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// MSSQLConfig represents the MSSQL connection configuration
type MSSQLConfig struct {
	ConnString string `json:"connString"`
}

// CheckMSSQLConnection checks the MSSQL connection using the provided connection string
func CheckMSSQLConnection(configFile string) error {
	// Read the configuration file
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Parse the JSON configuration
	var config map[string][]MSSQLConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return fmt.Errorf("failed to parse config JSON: %v", err)
	}

	// Get MSSQL configurations
	mssqlConfigs, ok := config["mssql"]
	if !ok || len(mssqlConfigs) == 0 {
		return fmt.Errorf("no MSSQL configurations found in config file")
	}

	// Check MSSQL connection for each configuration
	for _, mssqlConfig := range mssqlConfigs {
		db, err := sql.Open("sqlserver", mssqlConfig.ConnString)
		if err != nil {
			return fmt.Errorf("failed to open MSSQL connection: %v", err)
		}
		defer db.Close()

		if err := db.Ping(); err != nil {
			return fmt.Errorf("failed to ping MSSQL database: %v", err)
		}

		log.Printf("Successfully connected to MSSQL database")
	}

	return nil
}

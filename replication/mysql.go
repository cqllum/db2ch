package replication

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLConfig represents the MySQL connection configuration
type MySQLConfig struct {
	ConnString string `json:"connString"`
}

// CheckMySQLConnection checks the MySQL connection using the provided connection string
func CheckMySQLConnection(configFile string) error {
	// Read the configuration file
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Parse the JSON configuration
	var config map[string][]MySQLConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return fmt.Errorf("failed to parse config JSON: %v", err)
	}

	// Get MySQL configurations
	mysqlConfigs, ok := config["mysql"]
	if !ok || len(mysqlConfigs) == 0 {
		return fmt.Errorf("no MySQL configurations found in config file")
	}

	// Check MySQL connection for each configuration
	for _, mysqlConfig := range mysqlConfigs {
		db, err := sql.Open("mysql", mysqlConfig.ConnString)
		if err != nil {
			return fmt.Errorf("failed to open MySQL connection: %v", err)
		}
		defer db.Close()

		if err := db.Ping(); err != nil {
			return fmt.Errorf("failed to ping MySQL database: %v", err)
		}

		log.Printf("Successfully connected to MySQL database")
	}

	return nil
}

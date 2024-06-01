package replication

import (
	"log"
	"strings"
	"sync"

	"github.com/cqllum/db2ch/config"
)

var (
	running bool
	mu      sync.Mutex
)

// ReplicationService simulates replication service with database credentials
func ReplicationInit(dbType string, dbConfig config.DBConfig) error {
	// Parse values from dbConfig
	user := dbConfig.User
	password := dbConfig.Password
	dbName := dbConfig.DBName

	log.Printf("Replicating data for %s using credentials:\nUser: %s\nPassword: %s\nDBName: %s\n",
		dbType, user, password, dbName)

	switch strings.ToLower(dbType) {
	case "mysql":
		log.Println("ok")
	case "postgresql":
		log.Println("ok")
	case "mssql":
		log.Println("ok")
	default:
		log.Println("f")
	}

	return nil

	// Implement your replication logic using the database credentials securely

}

func StartReplication(dbType string, dbAlias string) {

	mu.Lock()
	defer mu.Unlock()
	// if running {
	// 	log.Println("Replication is already running@!")
	// 	return
	// }
	running = true

	config.LoadConfig()

	// Access the MySQL configuration with name "mysql1"
	var mysqlConfig *config.DBConfig
	for _, dbConfig := range config.AppConfig.Databases.MySQL {
		if dbConfig.Name == dbAlias {
			mysqlConfig = &dbConfig
			break
		}
	}

	if mysqlConfig != nil {
		log.Printf("Found MySQL configuration: %+v\n", *mysqlConfig)
	} else {
		log.Println("MySQL configuration with name 'mysql1' not found")
	}

}

func StopReplication() {
	mu.Lock()
	defer mu.Unlock()
	if !running {
		log.Println("Replication is not running")
		return
	}
	running = false
	log.Println("Stopping replication...")
	// Add logic to stop replication
}

func GetStatus() string {
	mu.Lock()
	defer mu.Unlock()
	if running {
		return "Replication is running"
	}
	return "Replication is stopped"
}

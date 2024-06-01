package replication

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/cqllum/db2ch/config"
)

var (
	running bool
	mu      sync.Mutex
)

// ReplicationService simulates replication service with database credentials
func ReplicationInit(dbType string, dbConfig config.DBConfig) error {
	// Parse values from dbConfig
	name := dbConfig.Name
	host := dbConfig.Host
	port := dbConfig.Port
	user := dbConfig.User
	password := dbConfig.Password
	dbName := dbConfig.DBName

	lockFilePath := filepath.Join(".", name+".lock")

	if _, err := os.Stat(lockFilePath); os.IsNotExist(err) {
		// Lock file doesn't exist, create it
		err := ioutil.WriteFile(lockFilePath, []byte{}, 0644)
		if err != nil {
			log.Fatalf("Error creating lock file: %v", err)
		}
		fmt.Println("Lock file created successfully.")

		stop := make(chan struct{})

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigCh
			log.Println("Received termination signal")
			close(stop) // Close the stop channel to signal the MySQL replication service to stop
		}()

		switch strings.ToLower(dbType) {
		case "mysql":
			log.Println("start mysql stuff")
			go MySQLReplicationService(name, host, int(port), user, password, dbName, stop)
		case "postgresql":
			log.Println("start psql stuff")
		case "mssql":
			log.Println("start mssql stuff")
		default:
			log.Println("nope")
		}

	} else {
		// Lock file already exists, indicate that the database is locked
		fmt.Println("Database is locked by another process.")
		return nil
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

package replication

import (
	"log"
	"sync"

	"github.com/cqllum/db2ch/config"
)

var (
	running bool
	mu      sync.Mutex
)

func StartReplication() {
	mu.Lock()
	defer mu.Unlock()
	// if running {
	// 	log.Println("Replication is already running@!")
	// 	return
	// }
	running = true
	log.Println("Starting replication...")

	for _, dbConfig := range config.AppConfig.MySQL {

		log.Println(dbConfig)
		//go replicateFromMySQL(dbConfig)

	}

	for _, dbConfig := range config.AppConfig.MSSQL {
		go replicateFromMSSQL(dbConfig)
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

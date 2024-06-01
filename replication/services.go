package replication

import (
	"log"
)

// ReplicationService is the interface for replication services.
type ReplicationService interface {
	Start() error
	Stop() error
}

// Start starts the MySQL replication service.
func MySQLReplicationService(name string, host string, port int, user string, password string, dbName string, stop chan struct{}) error {

	// Check if the database credentials work
	db, err := MySQLConnector(name, host, port, user, password)
	if err != nil {
		log.Printf("Error connecting to MySQL: %v", err)
		return err
	}
	defer db.Close()

	MySQLStartReplicator(name, host, uint16(port), user, password)
	return nil
}

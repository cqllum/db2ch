package replication

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnector(name string, host string, port int, user string, password string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close() // Close the connection before returning
		return nil, fmt.Errorf("error pinging MySQL: %v", err)
	}

	log.Println("Successfully connected to MySQL database")
	return db, nil
}

func MySQLStartReplicator(name string, host string, port uint16, user string, password string) error {

	// MySQL server configuration
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}

	// Create a binlog syncer
	syncer := replication.NewBinlogSyncer(cfg)

	// Start from the beginning of the binary log
	streamer, err := syncer.StartSync(mysql.Position{})
	if err != nil {
		log.Fatalf("Error starting binlog syncer: %v", err)
	}

	// Set up a signal handler to gracefully exit on Ctrl+C
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()

	// Map to store table names by table ID
	tableNames := make(map[uint64]string)

	// Process binlog events
	for {
		ev, err := streamer.GetEvent(ctx)
		if err != nil {
			if err == context.Canceled {
				log.Println("Binlog monitoring stopped")
				return nil
			}
			log.Printf("Error getting binlog event: %v", err)
			continue
		}

		// Process binlog event
		switch ev.Header.EventType {
		case replication.TABLE_MAP_EVENT:
			tableMapEvent := ev.Event.(*replication.TableMapEvent)
			tableNames[tableMapEvent.TableID] = string(tableMapEvent.Schema) + "." + string(tableMapEvent.Table)
		case replication.WRITE_ROWS_EVENTv1, replication.WRITE_ROWS_EVENTv2,
			replication.UPDATE_ROWS_EVENTv1, replication.UPDATE_ROWS_EVENTv2,
			replication.DELETE_ROWS_EVENTv1, replication.DELETE_ROWS_EVENTv2:
			rowsEvent := ev.Event.(*replication.RowsEvent)
			tableName := tableNames[rowsEvent.TableID]
			rows := make([][]interface{}, len(rowsEvent.Rows))
			for i, row := range rowsEvent.Rows {
				newRow := make([]interface{}, len(row))
				for j, val := range row {
					// Check if the value is a byte slice before conversion
					if byteVal, ok := val.([]byte); ok {
						newRow[j] = string(byteVal)
					} else {
						newRow[j] = val
					}
				}
				rows[i] = newRow
			}
			log.Printf("%s event: Table: %s, Rows: %+v", ev.Header.EventType, tableName, rows)
		}
	}

}

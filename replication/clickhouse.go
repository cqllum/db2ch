package replication

import (
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/cqllum/db2ch/config"
)

func ConnectClickHouse() *sql.DB {
	chConfig := config.AppConfig.ClickHouse
	dsn := "tcp://" + chConfig.Host + ":" + string(chConfig.Port) + "?username=" + chConfig.User + "&password=" + chConfig.Password + "&database=" + chConfig.DBName
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertDataToClickHouse(data []interface{}, dbConfig config.DBConfig) {
	db := ConnectClickHouse()
	defer db.Close()
	// Use dbConfig.ColumnMappings to map data types and insert data into ClickHouse
}

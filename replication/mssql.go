package replication

import (
	"database/sql"
	"log"

	"github.com/cqllum/db2ch/config"
	_ "github.com/denisenkom/go-mssqldb"
)

func replicateFromMSSQL(dbConfig config.DBConfig) {
	dsn := "sqlserver://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host + ":" + string(dbConfig.Port) + "?database=" + dbConfig.DBName
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Add logic to capture changes and use dbConfig.ColumnMappings
}

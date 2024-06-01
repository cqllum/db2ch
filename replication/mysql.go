package replication

import (
	"database/sql"
	"log"

	"github.com/cqllum/db2ch/config"
	_ "github.com/go-sql-driver/mysql"
)

/*
* Composed of several layers

[ Replication Slots ]
Get Replication Status
Create Replication Slot
Delete Replication Slot

[ Publication ]
Get Publication Status
Create Publication
Delete Publication

[ Table Publication ]
Get Publication Tables
Add Table to Publication
Remove Table to Publication


Remove Publication

*/

func replicateFromMySQL(dbConfig config.DBConfig) {
	log.Println("test")

	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + string(dbConfig.Port) + ")/" + dbConfig.DBName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Add logic to capture changes and use dbConfig.ColumnMappings
}

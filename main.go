package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/cqllum/db2ch/api"
	"github.com/cqllum/db2ch/replication"
)

type Config struct {
	Postgres []struct {
		ConnString string `json:"connString"`
	} `json:"postgres"`
	MySQL []struct {
		ConnString string `json:"connString"`
	} `json:"mysql"`
	MSSQL []struct {
		ConnString string `json:"connString"`
	} `json:"mssql"`
	ClickHouse struct {
		ConnString string `json:"connString"`
	} `json:"clickhouse"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var repServices []*replication.ReplicationService
	for _, pgConfig := range config.Postgres {
		repService, err := replication.NewReplicationService(pgConfig.ConnString)
		if err != nil {
			log.Printf("Failed to initialize PostgreSQL replication service: %v", err)
			continue
		}
		repServices = append(repServices, repService)
	}

	var mySQLRepServices []*replication.MySQLReplicationService
	for _, mySQLConfig := range config.MySQL {
		mySQLRepService, err := replication.NewMySQLReplicationService(mySQLConfig.ConnString)
		if err != nil {
			log.Printf("Failed to initialize MySQL replication service: %v", err)
			continue
		}
		mySQLRepServices = append(mySQLRepServices, mySQLRepService)
	}

	var mssqlRepServices []*replication.MSSQLReplicationService
	for _, mssqlConfig := range config.MSSQL {
		mssqlRepService, err := replication.NewMSSQLReplicationService(mssqlConfig.ConnString)
		if err != nil {
			log.Printf("Failed to initialize MSSQL replication service: %v", err)
			continue
		}
		mssqlRepServices = append(mssqlRepServices, mssqlRepService)
	}

	clickhouseService, err := replication.NewClickhouseService(config.ClickHouse.ConnString)
	if err != nil {
		log.Fatalf("Failed to initialize ClickHouse service: %v", err)
	}

	handler := api.NewHandler(repServices, mySQLRepServices, mssqlRepServices, clickhouseService)
	router := api.SetupRouter(handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

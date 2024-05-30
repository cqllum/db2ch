package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type DBConfig struct {
	Name           string            `json:"name"`
	Host           string            `json:"host"`
	Port           int               `json:"port"`
	User           string            `json:"user"`
	Password       string            `json:"password"`
	DBName         string            `json:"dbname"`
	ColumnMappings map[string]string `json:"column_mappings"`
}

type Config struct {
	PostgreSQL []DBConfig `json:"postgresql"`
	MySQL      []DBConfig `json:"mysql"`
	MSSQL      []DBConfig `json:"mssql"`
	ClickHouse DBConfig   `json:"clickhouse"`
}

var AppConfig Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}

//package main
package mySubLib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
	Log      bool
}

type DbClient struct {
	Db *sql.DB
}
func DbConnect(config *DbConfig) (*DbClient, error) {
	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.User, config.Pass, config.Host, config.Port, config.Database,
	)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	return &DbClient{Db: db}, nil
}

package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConfig struct {
	User     string
	Password string
	Database string
}

func Connect(config *SQLConfig) *sql.DB {
	connectionURI := fmt.Sprintf("%s:%s@/%s?parseTime=true", config.User, config.Password, config.Database)

	db, err := sql.Open("mysql", connectionURI)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

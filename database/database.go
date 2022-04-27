package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:docker@tcp(172.23.0.3:3306)/estacionamento")
	
	if err != nil {
		panic(err)
	}
	
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConectDB() *sql.DB {
	fmt.Println("Conectando ao banco de dados...")
	db, err := sql.Open("mysql", "root:docker@tcp(172.18.0.3:3306)/estacionamento")
	
	if err != nil {
		fmt.Println("Erro ao conectar")
		panic(err)
	}
	if err := db.Ping();err!=nil{
		fmt.Println("ping")
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
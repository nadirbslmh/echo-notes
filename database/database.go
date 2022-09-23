package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	DB_USERNAME string = "root"
	DB_PASSWORD string = ""
	DB_NAME     string = "echo_notes"
)

func Connect() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@/%s?parseTime=true",
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
	)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("error when creating a connection: %s", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")
}

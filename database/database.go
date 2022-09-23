package database

import (
	"database/sql"
	"echo-notes/util"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	DB_USERNAME string = util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = util.GetConfig("DB_PASSWORD")
	DB_NAME     string = util.GetConfig("DB_NAME")
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

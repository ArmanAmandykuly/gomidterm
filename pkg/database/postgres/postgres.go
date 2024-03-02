package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DBInit() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:lol@localhost/godb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}


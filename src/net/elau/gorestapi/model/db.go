package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSource string) {

	var err error

	db, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Panic(err)
	}
}

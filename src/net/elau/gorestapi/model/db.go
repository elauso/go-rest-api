package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(url string) {

	var err error

	db, err = sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
}

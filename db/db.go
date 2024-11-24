package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDB() *sql.DB {
	connection := "user=postgres dbname=alura_loja password=vini300695 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

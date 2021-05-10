package testhelpers

import (
	"bookstoreupdate/db"
	"database/sql"
	"fmt"
)

func ConnectDB() *db.DB {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "nhuttmt"
		dbname   = "bookstore"
	)
	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	//open db connection
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return &db.DB{
		Client: conn,
	}
}

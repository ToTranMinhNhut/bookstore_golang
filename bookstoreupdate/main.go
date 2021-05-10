package main

import (
	"log"
	"net/http"

	"bookstoreupdate/config"
	"bookstoreupdate/db"
	"bookstoreupdate/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars ", err)
	}

	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		log.Println("Connect to database is failded")
	}
	defer db.Close()

	routes := routes.RegisterRoutes(db)
	log.Println("starting server ")
	http.ListenAndServe(":8080", routes)
}

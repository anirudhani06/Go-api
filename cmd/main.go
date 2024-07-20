package main

import (
	"database/sql"
	"log"

	"github.com/anirudhani06/Go-api/cmd/api"
	"github.com/anirudhani06/Go-api/config"
	"github.com/anirudhani06/Go-api/db"
)

func main() {
	db, err := db.InitDb(config.Envs.DbAddr)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Db connected succesfully")
}

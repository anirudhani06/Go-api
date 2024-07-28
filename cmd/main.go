package main

import (
	"log"

	"github.com/anirudhani06/Go-api/cmd/api"
	"github.com/anirudhani06/Go-api/db"
)

func main() {

	store, err := db.PostgresStorage()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", store)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

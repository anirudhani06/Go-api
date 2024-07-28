package main

import (
	"log"

	"github.com/anirudhani06/Go-api/cmd/api"
)

func main() {

	server := api.NewAPIServer(":8080")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

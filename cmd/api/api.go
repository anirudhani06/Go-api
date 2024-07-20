package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/anirudhani06/Go-api/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}

}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userRoutes := user.NewHandler()
	userRoutes.UserRoutes(router)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	log.Println("Listening on: ", s.addr)
	return http.ListenAndServe(s.addr, router)

}

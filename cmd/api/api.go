package api

import (
	"log"
	"net/http"

	"github.com/anirudhani06/Go-api/db"
	"github.com/anirudhani06/Go-api/routes"
)

type APIServer struct {
	addr string
	db   *db.PostgresStore
}

func NewAPIServer(addr string, db *db.PostgresStore) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}

}

func (s *APIServer) Run() error {

	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	v1.Handle("/auth/", http.StripPrefix("/auth", routes.AuthRoutes()))
	v1.Handle("/users/", http.StripPrefix("/users", routes.UserRoutes()))

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	log.Println("Listening on: ", s.addr)
	return http.ListenAndServe(s.addr, mux)

}

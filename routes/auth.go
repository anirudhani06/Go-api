package routes

import (
	"net/http"

	"github.com/anirudhani06/Go-api/services"
)

func AuthRoutes() http.Handler {
	authMux := http.NewServeMux()

	authMux.HandleFunc("POST /login", services.HandleLogin)
	authMux.HandleFunc("POST /register", services.HandleRegister)
	return authMux
}

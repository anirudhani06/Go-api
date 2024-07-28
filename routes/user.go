package routes

import (
	"net/http"

	"github.com/anirudhani06/Go-api/services"
)

func UserRoutes() http.Handler {
	userMux := http.NewServeMux()
	userMux.HandleFunc("GET /", services.GetUser)
	return userMux
}

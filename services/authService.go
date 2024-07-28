package services

import (
	"net/http"

	"github.com/anirudhani06/Go-api/types"
	"github.com/anirudhani06/Go-api/utils"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJson(w, http.StatusCreated, payload)

}

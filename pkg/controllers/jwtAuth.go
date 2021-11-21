package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/models"
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/utils"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var UserCredentials *UserCredentials
	err := json.NewDecoder(r.Body).Decode(&UserCredentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authenticated, user := models.AuthenticateUser(UserCredentials.Email, UserCredentials.Password)
	if user.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User doesn't exists"))
		return
	}
	if !authenticated {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateJWT(UserCredentials.Email, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var response = map[string]interface{}{"token": token, "user": user}
	json.NewEncoder(w).Encode(response)
}

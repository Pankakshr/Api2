package handler

import (
	"encoding/json"
	"net/http"
	"vikash/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

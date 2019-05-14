package controller

import (
	"encoding/json"
	"go-api/domain"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(domain.GetUser())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(domain.GetUserID(params["id"]))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	domain.CreateUser(name, email)
	json.NewEncoder(w).Encode("User Created")
}

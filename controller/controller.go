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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	domain.UpdateUser(id, name, email)
	json.NewEncoder(w).Encode("User Updated")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	domain.DeleteUser(param["id"])
	json.NewEncoder(w).Encode("User Deleted")
}

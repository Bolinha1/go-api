package controller

import (
	"encoding/json"
	"go-api/domain"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(domain.GetUser())
}

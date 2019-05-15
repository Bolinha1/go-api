package router

import (
	"go-api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RoutesUsers() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

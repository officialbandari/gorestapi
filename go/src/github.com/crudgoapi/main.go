package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initialRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreatUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpadateUser).Methods("UPDATE")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	initialMigration()
	initialRouter()

}

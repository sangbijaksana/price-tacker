package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB("price-tracker.db")

	r := mux.NewRouter()

	r.HandleFunc("/register", Register).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(Authenticate)
	api.HandleFunc("/coins", GetTrackedCoins).Methods("GET")
	api.HandleFunc("/coins", AddCoin).Methods("POST")
	api.HandleFunc("/coins/{id}", RemoveCoin).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

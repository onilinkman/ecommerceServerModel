package main

import (
	"log"
	v1 "main/handlers/api/v1"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	createtables()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", v1.LoginUser).Methods("POST")
	/* r.Use(mux.CORSMethodMiddleware(r)) */
	log.Fatal(http.ListenAndServe(":8000", r))
}

func createtables() {
	models.CreateTableUsers()
}

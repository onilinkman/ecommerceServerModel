package main

import (
	"log"
	v1 "main/handlers/api/v1"
	"main/models"
	"main/templates"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	createtables()

	mux := mux.NewRouter()
	staticFiles := http.FileServer(http.Dir("build/static"))
	newTemp := templates.Temp{Path: "build/index.html"}
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))
	addUrlFrontLine(mux, newTemp)
	mux.HandleFunc("/api/v1/users", v1.LoginUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/session", v1.CheckUser)
	mux.HandleFunc("/api/v1/users/logout", v1.LogoutUser)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func createtables() {
	models.CreateTableUsers()
	models.CreateTableSession()
	models.SetGlobalEvent()
	models.CreateEventClearSession()
	models.CreateTableRol()
	models.CreateTableRol_user()
}

func addUrlFrontLine(mux *mux.Router, newTemp templates.Temp) {
	mux.HandleFunc("/", newTemp.RenderTemplate)
	mux.HandleFunc("/login", newTemp.RenderTemplate)
	mux.HandleFunc("/signup", newTemp.RenderTemplate)
	mux.HandleFunc("/addItem", newTemp.RenderTemplate)
}

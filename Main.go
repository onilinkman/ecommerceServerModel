package main

import (
	"fmt"
	"main/databasee"
	"main/models"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	databasee.CreateConnection()
	databasee.Ping()
	models.CreateTableUsers()
	fmt.Println(models.GetUser(1))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.ListenAndServe(":8080", nil)
}

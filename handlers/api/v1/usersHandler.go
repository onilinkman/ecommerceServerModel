package v1

import (
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	users, err := models.GetUserById(userId)

	if err != nil || len(*users) == 0 {
		models.SendNoContent(w)
	}

	models.SendData(w, users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	users, err := models.GetUserByEmail(email)
	if err != nil || len(*users) == 0 {
		models.SendNoContent(w)
	}
	models.SendData(w, users)
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := UserLogin{}
	fmt.Println(user, r.Body)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	}
	models.SendData(w, user)
}

package v1

import (
	"encoding/json"
	"main/models"
	"main/utils"
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

	if err != nil {
		models.SendNoContent(w)
	}

	models.SendData(w, users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	users, err := models.GetUserByEmail(email)
	if err != nil {
		models.SendNoContent(w)
	}
	models.SendData(w, users)
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Message struct {
	Message string `json:"message"`
	State   string `json:"state"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	userLogin := UserLogin{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&userLogin); err != nil {
		models.SendLoginFail(w, "Error al leer los Datos")
		return
	}

	user, err := models.GetUserByEmail(userLogin.Email)

	if err != nil {
		models.SendLoginFail(w, err.Error())
		return
	}

	if !user.ComparePassword(userLogin.Password) {
		models.SendLoginFail(w, "Usuario o Contraseña Incorrectos")
		return
	}

	cookie, expire := utils.SetSession(w)

	models.InsertSession(user.Id, cookie, expire)

	models.SendData(w, Message{Message: "Login Exitoso", State: "success"})
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("Cookie")
	models.DeleteSession(cookie)
	models.SendData(w, Message{Message: "Logout Exitoso", State: "success"})
}

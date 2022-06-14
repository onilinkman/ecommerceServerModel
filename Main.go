package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	v1 "main/handlers/api/v1"
	"main/models"
	"main/templates"
	"net/http"
	"net/smtp"
	"path/filepath"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	/*var mails []string
	mails = append(mails, "cmarban@umsa.bo")
	boo, err := SendEmailSMTP(mails, "hola", "email_template.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(boo)*/
	createtables()
	fmt.Println(`\`)

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

var emailAuth smtp.Auth

func SendEmailSMTP(to []string, data interface{}, template string) (bool, error) {
	emailHost := "smtp.gmail.com"
	emailFrom := "marbanchristian@gmail.com"
	emailPassword := "74094512ce"
	emailPort := 587

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	/*emailBody, err := parseTemplate(template, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}*/

	//mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	//subject := "Subject: " + "Test Email" + "!\n"
	//msg := []byte(subject + mime + "\n" + emailBody)
	msg := []byte("hola prueba")
	addr := fmt.Sprintf("%s:%d", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("gomail/email_templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
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

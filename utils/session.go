package utils

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	cookieName   = "go_session"
	cookieExpire = 24 * 2 * time.Hour
)

func SetSession(w http.ResponseWriter) (string, time.Time) {
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   uuid.NewV4().String(),
		Path:    "/",
		Expires: time.Now().Add(cookieExpire),
	}
	http.SetCookie(w, cookie)
	return cookie.Value, cookie.Expires
}

func DeleteSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

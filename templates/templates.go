package templates

import (
	"html/template"
	"log"
	"net/http"
)

type Temp struct {
	Path string
}

func (temp *Temp) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("build/index.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

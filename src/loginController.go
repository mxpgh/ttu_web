package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type loginController struct {
}

func (this *loginController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(path.Join(templatePath, "/html/login/index.html"))
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

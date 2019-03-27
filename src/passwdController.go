package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type Passwd struct {
	UserName string
	SysTime  string
}

type passwdController struct {
}

func (this *passwdController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles(path.Join(templatePath, "/html/passwd/index.html"))
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, &Passwd{user, getCurrentTime()})
}

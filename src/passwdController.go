package main

import (
	"html/template"
	"log"
	"net/http"
)

type Passwd struct {
	UserName string
	SysTime  string
}

type passwdController struct {
}

func (this *passwdController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles("template/html/passwd/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, &Passwd{user, getCurrentTime()})
}

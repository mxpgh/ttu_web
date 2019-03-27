package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type Upload struct {
	UserName string
	SysTime  string
	Token    string
}

type uploadController struct {
}

func (this *uploadController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles(path.Join(templatePath, "/html/upload/index.html"))
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, &Upload{
		user,
		getCurrentTime(),
		genToken()})
}

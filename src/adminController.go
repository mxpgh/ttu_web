package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type User struct {
	UserName string
	SysTime  string
}

type adminController struct {
}

func (this *adminController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	//fp := path.Join(templatePath, "/html/admin/index.html")
	//log.Println(templatePath)
	//log.Println(fp)
	t, err := template.ParseFiles(path.Join(templatePath, "/html/admin/index.html"))
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, &User{user, getCurrentTime()})
}

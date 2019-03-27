package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type SysMgr struct {
	UserName string
	SysTime  string
}

type sysmgrController struct {
}

func (this *sysmgrController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles(path.Join(templatePath, "/html/sysmgr/index.html"))
	if err != nil {
		log.Println(err)
		return
	}

	t.Execute(w, &SysMgr{
		user,
		getCurrentTime()})
}

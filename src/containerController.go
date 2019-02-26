package main

import (
	"html/template"
	"log"
	"net/http"
)

type CTItem struct {
	Name    string
	CPURate string
	MemRate string
}

type Container struct {
	UserName string
	SysTime  string
	CTList   []CTItem
}

type containerController struct {
}

func (this *containerController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles("template/html/container/index.html")
	if err != nil {
		log.Println(err)
	}

	var ctList []CTItem
	gCLRW.RLock()
	for i, v := range gContainerList {
		ct := CTItem{}
		ct.Name = v.Name
		ct.CPURate = v.CPU
		ct.MemRate = v.Memory
		ctList = append(ctList, ct)
		log.Println(i, " ", v.Container, " ", v.Name, " ", v.CPU, " ", v.Memory)
	}
	gCLRW.RUnlock()

	t.Execute(w, &Container{
		user,
		getCurrentTime(),
		ctList})
}

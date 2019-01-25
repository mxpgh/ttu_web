package main

import (
	"html/template"
	"log"
	"net/http"
)

type Status struct {
	UserName                   string
	DevCurTime                 string
	DevTemp                    string
	OsCPURate                  string
	OsMemRate                  string
	OsDiskRate                 string
	ContainerCPURate           string
	ContainerMemRate           string
	AppCPURate                 string
	AppMemRate                 string
	RTCFault                   string
	TempDetectFault            string
	PeripheralHardwareFault    string
	ContainerList              string
	CommunicationNetworkStatus string
}

type statusController struct {
}

func (this *statusController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles("template/html/status/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, &Status{
		user,
		getDevCurrentTime(),
		getDevTemperature(),
		getOsCPURate(),
		getOsMemoryRate(),
		getOsDiskRate(),
		getContainerCPURate(),
		getContainerMemoryRate(),
		getAppCPURate(),
		getAppMemoryRate(),
		getRTCFault(),
		getTemperatureDetectFault(),
		getPeripheralHardwareFault(),
		getContainerList(),
		getCommunicationNetworkStatus()})
}

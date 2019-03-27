package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type Status struct {
	UserName                   string
	SysTime                    string
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
	t, err := template.ParseFiles(path.Join(templatePath, "/html/status/index.html"))
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, &Status{
		user,
		getCurrentTime(),
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

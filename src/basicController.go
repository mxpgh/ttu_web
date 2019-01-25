package main

import (
	"html/template"
	"log"
	"net/http"
)

type Basic struct {
	UserName         string
	DevType          string
	DevName          string
	DevLabel         string
	Vendor           string
	DevStatus        string
	DevMac           string
	DevCurTime       string
	DevStartTime     string
	DevRunTimes      string
	DevMem           string
	DevDisk          string
	DevSoftPatch     string
	AppPatch         string
	Hardware         string
	DevCommunication string
	Platform         string
	K8sInfo          string
	DockerInfo       string
}

type basicController struct {
}

func (this *basicController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles("template/html/basic/index.html")
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, &Basic{user,
		getDevType(),
		getDevName(),
		getDevLabel(),
		getVendor(),
		getDevStatus(),
		getDevMac(),
		getDevCurrentTime(),
		getDevStartTime(),
		getDevRunTimes(),
		getDevMemory(),
		getDevDisk(),
		getSoftPatch(),
		getAppPatch(),
		getHardwareVer(),
		getCommunicationInterface(),
		getPlatformInfo(),
		getK8sInfo(),
		getDockerInfo()})

}

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Config struct {
	UserName                string
	SysTime                 string
	MainStationIPv4         string
	MainStationIPv4Port     string
	MainStationIPv6         string
	MainStationIPv6Port     string
	BackMainStationIPv4     string
	BackMainStationIPv4Port string
	BackMainStationIPv6     string
	BackMainStationIPv6Port string
	SysCPURateUpper         string
	SysMemRateUpper         string
	SysDiskRateUpper        string
	SysMonitorWndTime       string
	ContainerCPURateUpper   string
	ContainerMemRateUpper   string
	ContainerMonitorWndTime string
	AppCPURateUpper         string
	AppMemRateUpper         string
	AppMonitorWndTime       string
	TempLower               string
	TempUpper               string
	TempUpperWnd            string
}

type configController struct {
}

func (this *configController) IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles("template/html/config/index.html")
	if err != nil {
		log.Println(err)
	}
	tempLower, tempUpper := getTemperatureUpper()
	t.Execute(w, &Config{
		user,
		getCurrentTime(),
		getMainStationIPv4(),
		getMainStationIPv4Port(),
		getMainStationIPv6(),
		getMainStationIPv6Port(),
		getBackMainStationIPv4(),
		getBackMainStationIPv4Port(),
		getBackMainStationIPv6(),
		getBackMainStationIPv6Port(),
		getSysCPURateUpper(),
		getSysMemoryRateUpper(),
		getSysDiskRateUpper(),
		getSysMonitorWndTime(),
		getContainerCPURateUpper(),
		getContainerMemoryRateUpper(),
		getContainerMonitorWndTime(),
		getAppCPURateUpper(),
		getAppMemoryRateUpper(),
		getAppMonitorWndTime(),
		tempLower,
		tempUpper,
		getTemperatureUpperWnd()})

}

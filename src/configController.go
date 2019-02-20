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
	MainStationIPv4Port     uint16
	MainStationIPv6         string
	MainStationIPv6Port     uint16
	BackMainStationIPv4     string
	BackMainStationIPv4Port uint16
	BackMainStationIPv6     string
	BackMainStationIPv6Port uint16
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

	tempLower, tempUpper := getTemperatureThreshold()
	ipv4, ipv4Port := getMainStationIPv4()
	ipv6, ipv6Port := getMainStationIPv6()
	backIpv4, backIpv4Port := getBackMainStationIPv4()
	backIpv6, backIpv6Port := getBackMainStationIPv6()

	t.Execute(w, &Config{
		user,
		getCurrentTime(),
		ipv4,
		uint16(ipv4Port),
		ipv6,
		uint16(ipv6Port),
		backIpv4,
		uint16(backIpv4Port),
		backIpv6,
		uint16(backIpv6Port),
		getSysCPURateThreshold(),
		getSysMemoryRateThreshold(),
		getSysDiskRateThreshold(),
		getSysMonitorWndTime(),
		getContainerCPURateThreshold(),
		getContainerMemoryRateThreshold(),
		getContainerMonitorWndTime(),
		getAppCPURateThreshold(),
		getAppMemoryRateThreshold(),
		getAppMonitorWndTime(),
		tempLower,
		tempUpper,
		getTemperatureThresholdWnd()})

}

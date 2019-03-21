package main

import (
	"html/template"
	"log"
	"net/http"
)

type Config struct {
	UserName string
	SysTime  string

	MainStationIPv4         string
	MainStationIPv4Port     uint16
	MainStationIPv6         string
	MainStationIPv6Port     uint16
	BackMainStationIPv4     string
	BackMainStationIPv4Port uint16
	BackMainStationIPv6     string
	BackMainStationIPv6Port uint16

	OPSMainStationIPv4         string
	OPSMainStationIPv4Port     uint16
	OPSMainStationIPv6         string
	OPSMainStationIPv6Port     uint16
	BackOPSMainStationIPv4     string
	BackOPSMainStationIPv4Port uint16
	BackOPSMainStationIPv6     string
	BackOPSMainStationIPv6Port uint16

	LocalIPv4        string
	LocalIPv6        string
	LocalIPv4Route   string
	LocalIPv6Route   string
	LocalIPv4SubMask string
	LocalIPv6SubMask string
	LocalIPv4DNS     string
	LocalIPv6DNS     string

	SysCPURateUpper   string
	SysMemRateUpper   string
	SysDiskRateUpper  string
	SysMonitorWndTime string

	ContainerCPURateUpper   string
	ContainerMemRateUpper   string
	ContainerMonitorWndTime string

	AppCPURateUpper   string
	AppMemRateUpper   string
	AppMonitorWndTime string

	TempLower    string
	TempUpper    string
	TempUpperWnd string
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

	opsIpv4, opsIpv4Port := getOPSMainStationIPv4()
	opsIpv6, opsIpv6Port := getOPSMainStationIPv6()
	backOPSIpv4, backOPSIpv4Port := getBackOPSMainStationIPv4()
	backOPSIpv6, backOPSIpv6Port := getBackOPSMainStationIPv6()

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

		opsIpv4,
		uint16(opsIpv4Port),
		opsIpv6,
		uint16(opsIpv6Port),
		backOPSIpv4,
		uint16(backOPSIpv4Port),
		backOPSIpv6,
		uint16(backOPSIpv6Port),

		getLocalIPv4(),
		getLocalIPv6(),
		getLocalIPv4Route(),
		getLocalIPv6Route(),
		getLocalIPv4SubMask(),
		getLocalIPv6SubMask(),
		getLocalIPv4DNS(),
		getLocalIPv6DNS(),

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

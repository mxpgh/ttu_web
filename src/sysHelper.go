// +build windows

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	container = iota
	app
)

const (
	cpu = iota
	memory
)

var (
	gk8sVer           string
	gContaninerCPU    string
	gContaninerMemory string
	gCLRW             sync.RWMutex
	gContainerList    []dockerStat
)

type dockerStat struct {
	Container string
	Name      string
	Memory    string
	CPU       string
}

func timeTask() {
	var waitConfTime = 5 * time.Second
	var waitTaskTime = 5 * time.Second
	var waitContainerTime time.Duration
	var waitAppTime time.Duration
	var containerCPURateThreshold, containerMemRateThreshold, AppCPURateThreshold, AppMemRateThreshold int
	startConfTime := time.Now().UTC()
	startTaskTime := time.Now().UTC()
	startContainerTime := time.Now().UTC()
	startAppTime := time.Now().UTC()

	for {
		endTime := time.Now().UTC()
		var durationConf = endTime.Sub(startConfTime)
		if durationConf >= waitConfTime {
			waitContainerTime = time.Duration(getContainerMonitorWndIntTime()) * time.Minute
			waitAppTime = time.Duration(getAppMonitorWndIntTime()) * time.Minute
			containerCPURateThreshold = getContainerCPURateIntThreshold()
			containerMemRateThreshold = getContainerMemoryRateIntThreshold()
			AppCPURateThreshold = getAppCPURateIntThreshold()
			AppMemRateThreshold = getAppMemoryRateIntThreshold()

			startConfTime = time.Now().UTC()
			log.Printf("waitConf %v %v %.3f", waitConfTime, durationConf, durationConf.Seconds())
		}

		endTime = time.Now().UTC()
		var durationTask = endTime.Sub(startTaskTime)
		if durationTask >= waitTaskTime {
			gk8sVer = execBashCmd("kubelet --version")

			var cl []dockerStat
			gContaninerCPU, gContaninerMemory, cl = getDockerStat()

			gCLRW.Lock()
			gContainerList = gContainerList[:0]
			gContainerList = append(gContainerList, cl...)
			gCLRW.Unlock()

			startTaskTime = time.Now().UTC()
			log.Printf("waitTask %v %v %.3f", waitTaskTime, durationTask, durationTask.Seconds())

			var durationContainer = endTime.Sub(startContainerTime)
			if durationContainer >= waitContainerTime {
				for _, v := range cl {
					f, err := strconv.ParseFloat(strings.TrimRight(v.CPU, "%"), 32)
					if err != nil {
						continue
					}
					if int(f*100) > containerCPURateThreshold {

					}

					f, err = strconv.ParseFloat(strings.TrimRight(v.Memory, "%"), 32)
					if err != nil {
						continue
					}
					if int(f*100) > containerMemRateThreshold {

					}
				}

				startContainerTime = time.Now().UTC()
				log.Printf("waitContainer %v %v %.3f", waitContainerTime, durationContainer, durationContainer.Seconds())
			}

			var durationApp = endTime.Sub(startAppTime)
			if durationApp >= waitAppTime {
				_ = AppCPURateThreshold
				_ = AppMemRateThreshold

				startAppTime = time.Now().UTC()
				log.Printf("waitApp %v %v %.3f", waitAppTime, durationApp, durationApp.Seconds())
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func reboot() error {
	return nil
}

func getDevType() string {
	return "TTU"
}

func getDevName() string {
	return "TTU0001"
}

func getDevLabel() string {
	return "TTU_SZBC0001"
}

func getVendor() string {
	return "Step electronics"
}

func getDevStatus() string {
	return "正常状态"
}

func getDevMac() string {
	return "00:14:97:22:8c:4e"
}

func getDevCurrentTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func getDevStartTime() string {
	return execBashCmd(`date -d "$(awk -F. '{print $1}' /proc/uptime) second ago" +"%Y-%m-%d %H:%M:%S"`)
}

func getDevRunTimes() string {
	return execBashCmd(`awk -F. '{print $1}' /proc/uptime`)
}

func getDevMemory() string {
	return execBashCmd(`free -m |grep "Mem:" | awk '{print $2}'`) + "M"
}

func getDevDisk() string {
	return execBashCmd(`df -h / | awk '{print $2}' | sed -n '2p'`)
}

func getSoftPatch() string {
	return "V1.0.0.0"
}

func getAppPatch() string {
	return "v1.0.0.0"
}

func getHardwareVer() string {
	return "v1.0.0.0"
}

func getCommunicationInterface() string {
	return ""
}

func getPlatformInfo() string {
	return execBashCmd(`uname -a`)
}

func getK8sInfo() string {
	return gk8sVer
}

func getDockerInfo() string {
	return execBashCmd(`docker -v | awk '{print $3}'`)
}

func getDevTemperature() string {
	return "10"
}

func getOsCPURate() string {
	return "0.3"
}

func getOsMemoryRate() string {
	return "0.5"
}

func getOsDiskRate() string {
	return "0.6"
}

func getContainerCPURate() string {
	return gContaninerCPU
}

func getContainerMemoryRate() string {
	return gContaninerMemory
}

func getAppCPURate() string {
	return "0.1"
}

func getAppMemoryRate() string {
	return "0.2"
}

func getRTCFault() string {
	return ""
}

func getTemperatureDetectFault() string {
	return ""
}

func getPeripheralHardwareFault() string {
	return ""
}

func getContainerList() string {
	return ""
}

func getCommunicationNetworkStatus() string {
	return ""
}

func getDockerStat() (cpu, mem string, dL []dockerStat) {
	ret := execBashCmd(`docker stats --no-stream --format \
	"{\"container\":\"{{ .Container }}\",\"name\":\"{{ .Name }}\",\"memory\":\"{{ .MemPerc }}\",\"cpu\":\"{{ .CPUPerc }}\"}"`)
	//log.Println(ret)

	totMem := 0.0
	totCPU := 0.0
	trArr := strings.Split(ret, "\n")
	for _, v := range trArr {
		if v == "" {
			continue
		}
		//log.Println(v)
		vb := []byte(v)
		//log.Println(vb)
		ds := dockerStat{}
		err := json.Unmarshal(vb, &ds)
		if err != nil {
			log.Println(err)
			//log.Println(vb)
		} else {
			//log.Println("memory: ", ds.Memory, "cpu: ", ds.Cpu)
			//log.Println(strings.TrimRight(ds.Memory, "%"))
			dL = append(dL, ds)
			f, err := strconv.ParseFloat(strings.TrimRight(ds.Memory, "%"), 32)
			if err == nil {
				totMem += f
			}
			f, err = strconv.ParseFloat(strings.TrimRight(ds.CPU, "%"), 32)
			if err == nil {
				totCPU += f
			}
		}
	}
	mem = fmt.Sprintf("%.2f", totMem)
	cpu = fmt.Sprintf("%.2f", totCPU)
	//log.Println("docker mem: ", fmt.Sprintf("%.2f", totMem), ", cpu: ", fmt.Sprintf("%.2f", totCpu))
	return
}

//////////////////////////////////////////
func setMainStationIPv4(ip string, port string) error {
	return nil
}

func getMainStationIPv4() (ip string, port uint16) {
	buf := make([]byte, 16)
	log.Println(buf)
	log.Println(byteToString(buf))
	return byteToString(buf), 6443
}

func setBackMainStationIPv4(ip string, port string) error {
	return nil
}

func getBackMainStationIPv4() (ip string, port uint16) {
	return "192.168.1.0", 6443
}

///////////////////////////////////////////
func setMainStationIPv6(ip string, port string) error {
	return nil
}

func getMainStationIPv6() (ip string, port uint16) {
	return "fe80::c10c:9d86:382f:4797", 6443
}

func setBackMainStationIPv6(ip string, port string) error {
	return nil
}

func getBackMainStationIPv6() (ip string, port uint16) {
	return "fe80::c10c:9d86:382f:4797", 6443
}

//////////////////////////////////////////
func setOPSMainStationIPv4(ip string, port string) error {
	return nil
}

func getOPSMainStationIPv4() (ip string, port uint16) {
	buf := make([]byte, 16)
	log.Println(buf)
	log.Println(byteToString(buf))
	return byteToString(buf), 6443
}

func setBackOPSMainStationIPv4(ip string, port string) error {
	return nil
}

func getBackOPSMainStationIPv4() (ip string, port uint16) {
	return "192.168.1.0", 6443
}

///////////////////////////////////////////
func setOPSMainStationIPv6(ip string, port string) error {
	return nil
}

func getOPSMainStationIPv6() (ip string, port uint16) {
	return "fe80::c10c:9d86:382f:4797", 6443
}

func setBackOPSMainStationIPv6(ip string, port string) error {
	return nil
}

func getBackOPSMainStationIPv6() (ip string, port uint16) {
	return "fe80::c10c:9d86:382f:4797", 6443
}

/////////////////////////////////////////////////////////////////////////////////////
func getLocalIPv4() string {
	return "192.168.1.12"
}

func setLocalIPv4(ip string) error {
	return nil
}

func getLocalIPv6() string {
	return "fe80::c10c:9d86:382f:4797"
}

func setLocalIPv6(ip string) error {
	return nil
}

func getLocalIPv4Route() string {
	return "192.168.1.1"
}

func setLocalIPv4Route(route string) error {
	return nil
}

func getLocalIPv6Route() string {
	return "fe80::c10c:9d86:382f:4797"
}

func setLocalIPv6Route(route string) error {
	return nil
}

func getLocalIPv4SubMask() string {
	return "255.255.255.0"
}

func setLocalIPv4SubMask(mask string) error {
	return nil
}

func getLocalIPv6SubMask() string {
	return "fe80::c10c:9d86:382f:4797"
}

func setLocalIPv6SubMask(mask string) error {
	return nil
}

func getLocalIPv4DNS() string {
	return "192.168.1.1"
}

func setLocalIPv4DNS(dns string) error {
	return nil
}

func getLocalIPv6DNS() string {
	return "fe80::c10c:9d86:382f:4797"
}

func setLocalIPv6DNS(dns string) error {
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////
func setSysCPURateThreshold(rate string) error {
	return nil
}

func getSysCPURateThreshold() string {
	return "80"
}

func setSysMemoryRateThreshold(rate string) error {
	return nil
}

func getSysMemoryRateThreshold() string {
	return "80"
}

func setSysDiskRateThreshold(rate string) error {
	return nil
}

func getSysDiskRateThreshold() string {
	return "80"
}

///////////////////////////////////////////////
func setSysMonitorWndTime(wnd string) error {
	return nil
}

func getSysMonitorWndTime() string {
	return "10"
}

func setContainerCPURateThreshold(rate string) error {
	return nil
}

func getContainerCPURateThreshold() string {
	return "80"
}

func getContainerCPURateIntThreshold() int {
	return 80
}

func setContainerMemoryRateThreshold(rate string) error {
	return nil
}

func getContainerMemoryRateThreshold() string {
	return "80"
}

func getContainerMemoryRateIntThreshold() int {
	return 80
}

// 单位：分钟
func setContainerMonitorWndTime(wnd string) error {
	return nil
}

// 单位：分钟
func getContainerMonitorWndTime() string {
	return "5"
}

// 单位：分钟
func getContainerMonitorWndIntTime() int {
	return 5
}

/////////////////////////////////////////////
func setAppCPURateThreshold(rate string) error {
	return nil
}

func getAppCPURateThreshold() string {
	return "80"
}

func getAppCPURateIntThreshold() int {
	return 80
}

func setAppMemoryRateThreshold(rate string) error {
	return nil
}

func getAppMemoryRateThreshold() string {
	return "80"
}

func getAppMemoryRateIntThreshold() int {
	return 80
}

// 单位：分钟
func setAppMonitorWndTime(wnd string) error {
	return nil
}

// 单位：分钟
func getAppMonitorWndTime() string {
	return "5"
}

// 单位：分钟
func getAppMonitorWndIntTime() int {
	return 5
}

/////////////////////////////////////////////
func setTemperatureThreshold(lower, upper string) error {
	return nil
}

func getTemperatureThreshold() (lower, upper string) {
	return "-40", "40"
}

func setTemperatureThresholdWnd(wnd string) error {
	return nil
}

func getTemperatureThresholdWnd() string {
	return "10"
}

// 告警通知
func pushAlarm(appType, resType int, name string, value int) error {
	return nil
}

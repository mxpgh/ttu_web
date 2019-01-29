package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	gk8sVer           string
	gContaninerCPU    string
	gContaninerMemory string
)

type dockerStat struct {
	Container string
	Memory    string
	CPU       string
}

func timeTask() {
	for {
		gk8sVer = execBashCmd("kubelet --version")
		gContaninerCPU, gContaninerMemory = getDockerStat()
		time.Sleep(5 * time.Second)
	}
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

func getDockerStat() (cpu, mem string) {
	ret := execBashCmd(`docker stats --no-stream --format \
	"{\"container\":\"{{ .Container }}\",\"memory\":\"{{ .MemPerc }}\",\"cpu\":\"{{ .CPUPerc }}\"}"`)
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
func setMainStationIPv4(ip string) error {
	return nil
}

func getMainStationIPv4() string {
	return "192.168.1.0"
}

func setMainStationIPv4Port(port string) error {
	return nil
}

func getMainStationIPv4Port() string {
	return "6443"
}

func setBackMainStationIPv4(ip string) error {
	return nil
}

func getBackMainStationIPv4() string {
	return "192.168.1.0"
}

func setBackMainStationIPv4Port(port string) error {
	return nil
}

func getBackMainStationIPv4Port() string {
	return "6443"
}

///////////////////////////////////////////
func setMainStationIPv6(ip string) error {
	return nil
}

func getMainStationIPv6() string {
	return "fe80::c10c:9d86:382f:4797"
}

func setMainStationIPv6Port(port string) error {
	return nil
}

func getMainStationIPv6Port() string {
	return "6443"
}

func setBackMainStationIPv6(ip string) error {
	return nil
}

func getBackMainStationIPv6() string {
	return "fe80::c10c:9d86:382f:4797"
}

func setBackMainStationIPv6Port(port string) error {
	return nil
}

func getBackMainStationIPv6Port() string {
	return "6443"
}

////////////////////////////////////////////
func setSysCPURateUpper(rate string) error {
	return nil
}

func getSysCPURateUpper() string {
	return "80"
}

func setSysMemoryRateUpper(rate string) error {
	return nil
}

func getSysMemoryRateUpper() string {
	return "80"
}

func setSysDiskRateUpper(rate string) error {
	return nil
}

func getSysDiskRateUpper() string {
	return "80"
}

///////////////////////////////////////////////
func setSysMonitorWndTime(wnd string) error {
	return nil
}

func getSysMonitorWndTime() string {
	return "10"
}

func setContainerCPURateUpper(rate string) error {
	return nil
}

func getContainerCPURateUpper() string {
	return "80"
}

func setContainerMemoryRateUpper(rate string) error {
	return nil
}

func getContainerMemoryRateUpper() string {
	return "80"
}

func setContainerMonitorWndTime(wnd string) error {
	return nil
}

func getContainerMonitorWndTime() string {
	return "5"
}

/////////////////////////////////////////////
func setAppCPURateUpper(rate string) error {
	return nil
}

func getAppCPURateUpper() string {
	return "80"
}

func setAppMemoryRateUpper(rate string) error {
	return nil
}

func getAppMemoryRateUpper() string {
	return "80"
}

func setAppMonitorWndTime(wnd string) error {
	return nil
}

func getAppMonitorWndTime() string {
	return "5"
}

/////////////////////////////////////////////
func setTemperatureUpper(lower, upper string) error {
	return nil
}

func getTemperatureUpper() (lower, upper string) {
	return "-40", "40"
}

func setTemperatureUpperWnd(wnd string) error {
	return nil
}

func getTemperatureUpperWnd() string {
	return "10"
}

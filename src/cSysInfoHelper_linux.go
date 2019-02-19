package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lsysconfig -lyaml-cpp
#include <stdlib.h>
#include "sysconfig.h"
*/
import "C"

var (
	gCPURate   string
	gk8sVer    string
	gAppCPU    string
	gAppMemory string
)

type dockerStat struct {
	Container string
	Memory    string
	CPU       string
}

func timeTask() {
	for {
		gk8sVer = execBashCmd("kubelet --version")
		gAppCPU, gAppMemory = getDockerStat()

		{
			rate := C.getCpuOccupy()
			gCPURate = strconv.Itoa(int(rate)) + "%"
		}
		time.Sleep(5 * time.Second)
	}
}

func getDevType() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getDevType((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getDevName() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getDevName((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getDevLabel() string {
	return "TTU_SZBC0001"
}

func getVendor() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getDevVendor((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getDevStatus() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getDevStatus((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getDevMac() string {
	return "00:14:97:22:8c:4e"
}

func getDevCurrentTime() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getTime((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
	//return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func getDevStartTime() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getUPTime((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getDevRunTimes() string {
	times := C.getRunTime()
	return fmtTimes(int(times))
	//return strconv.Itoa(int(times)) + " s"
	//return execBashCmd(`awk -F. '{print $1}' /proc/uptime`)
}

func getDevMemory() string {
	mem := C.getRamSize()
	return strconv.Itoa(int(mem)) + " KB"
	//return execBashCmd(`free -m |grep "Mem:" | awk '{print $2}'`) + "M"
}

func getDevDisk() string {
	disk := C.getDiskSize()
	return strconv.FormatFloat(float64(disk), 'f', 2, 64) + " GB"
	//return execBashCmd(`df -h / | awk '{print $2}' | sed -n '2p'`)
}

func getSoftPatch() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getSoftwareVer((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getAppPatch() string {
	return "v1.0.0.0"
}

func getHardwareVer() string {
	inlen := C.int(128)
	buf := make([]byte, 128)
	outlen := C.getHardwareVer((*C.char)(unsafe.Pointer(&buf[0])), inlen)
	return string(buf[:outlen])
}

func getCommunicationInterface() string {
	return ""
}

func getPlatformInfo() string {
	return execBashCmd(`uname -oi`)
}

func getK8sInfo() string {
	return gk8sVer
}

func getDockerInfo() string {
	return execBashCmd(`docker -v | awk '{print $3}' | awk '{split($0,b,",");print b[1]}'`)
}

func getDevTemperature() string {
	temp := C.getTemperature()
	return strconv.Itoa(int(temp)) + " ℃"
}

func getOsCPURate() string {
	return gCPURate
	//rate := C.getCpuOccupy()
	//return strconv.Itoa(int(rate)) + "%"
}

func getOsMemoryRate() string {
	rate := C.getRamOccupy()
	return strconv.Itoa(int(rate)) + "%"
}

func getOsDiskRate() string {
	rate := C.getDiskOccupy()
	return strconv.Itoa(int(rate)) + "%"
}

func getContainerCPURate() string {
	return execBashCmd(`ps -aux | grep dockerd | grep -v grep | awk '{print$3}'`) + "%"
}

func getContainerMemoryRate() string {
	return execBashCmd(`ps -aux | grep dockerd | grep -v grep | awk '{print$4}'`) + "%"
}

func getAppCPURate() string {
	return gAppCPU + "%"
}

func getAppMemoryRate() string {
	return gAppMemory + "%"
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
//set system current time
func setSysCurTime(tm string) error {
	cs := C.CString(tm)
	ret := C.setTime(cs)
	_ = ret
	C.free(unsafe.Pointer(cs))
	return nil
}

func setSysCPURateUpper(rate string) error {
	upper, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}
	ret := C.setCpuThreshold(C.int(upper))
	_ = ret
	return nil
}

func getSysCPURateUpper() string {
	upper := C.int(0)
	ret := C.getCpuThreshold((*C.int)(unsafe.Pointer(&upper)))
	_ = ret
	return strconv.Itoa(int(upper))
}

func setSysMemoryRateUpper(rate string) error {
	upper, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}
	ret := C.setRamThreshold(C.int(upper))
	_ = ret
	return nil
}

func getSysMemoryRateUpper() string {
	upper := C.int(0)
	ret := C.getRamThreshold((*C.int)(unsafe.Pointer(&upper)))
	_ = ret
	return strconv.Itoa(int(upper))
}

func setSysDiskRateUpper(rate string) error {
	upper, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}
	ret := C.setDiskThreshold(C.int(upper))
	_ = ret
	return nil
}

func getSysDiskRateUpper() string {
	upper := C.int(0)
	ret := C.getDiskThreshold((*C.int)(unsafe.Pointer(&upper)))
	_ = ret
	return strconv.Itoa(int(upper))
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
	tempUpper, err := strconv.Atoi(upper)
	if err != nil {
		return err
	}
	tempLower, err := strconv.Atoi(lower)
	if err != nil {
		return err
	}
	ret := C.setTempThreshold(C.int(tempUpper), C.int(tempLower))
	_ = ret
	return nil
}

func getTemperatureUpper() (lower, upper string) {
	tempUpper := C.int(0)
	tempLower := C.int(0)
	ret := C.getTempThreshold((*C.int)(unsafe.Pointer(&tempUpper)), (*C.int)(unsafe.Pointer(&tempLower)))
	_ = ret
	return strconv.Itoa(int(tempLower)), strconv.Itoa(int(tempUpper))
}

func setTemperatureUpperWnd(wnd string) error {
	return nil
}

func getTemperatureUpperWnd() string {
	return "10"
}

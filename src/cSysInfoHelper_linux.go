package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unsafe"
	"sync"
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
	gCLRW	sync.RWMutex
	gContainerList []dockerStat
)

type dockerStat struct {
	Container string
	Name	  string
	Memory    string
	CPU       string
}

func timeTask() {
	for {
		gk8sVer = execBashCmd("kubelet --version")
		var cl []dockerStat
		gAppCPU, gAppMemory, cl = getDockerStat()
		gCLRW.Lock()
		gContainerList = gContainerList[:0]
		gContainerList = append(gContainerList, cl...)
		gCLRW.Unlock()

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
	//mem := C.getRamSize()
	//return strconv.Itoa(int(mem)) + " KB"
	return execBashCmd(`free -h |grep "Mem:" | awk '{print $2}'`) + "B"
}

func getDevDisk() string {
	disk := C.getDiskSize()
	return strconv.FormatFloat(float64(disk), 'f', 2, 64) + "GB"
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
	temp := C.int(0)
	ret := C.getTemperature(&temp)
	if 0 == ret {
		return "传感器故障"
	}
	return strconv.Itoa(int(temp)) + "℃"
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
	tm := time.Now().UnixNano() / int64(time.Millisecond)
	ret := C.getRtcStatus()
	ed := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("getRtcStatus time: ", ed-tm)
	if 0 == ret {
		return "异常"
	}
	return "正常"
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
	cs := C.CString(ip)
	pt, _ := strconv.Atoi(port)
	ret := C.setServer(C.int(0), C.int(4), cs, C.int(pt))
	C.free(unsafe.Pointer(cs))
	if 0 == ret {
		return errors.New("set main station ipv4 error")
	}
	return nil
}

func getMainStationIPv4() (ip string, port uint16) {
	cPort := C.int(0)
	cIp := make([]byte, 16)
	ret := C.getServer(C.int(0), C.int(4), (*C.char)(unsafe.Pointer(&cIp[0])), &cPort)
	_ = ret
	return byteToString(cIp), uint16(cPort)
}

func setBackMainStationIPv4(ip string, port string) error {
	cs := C.CString(ip)
	pt, _ := strconv.Atoi(port)
	ret := C.setServer(C.int(1), C.int(4), cs, C.int(pt))
	C.free(unsafe.Pointer(cs))
	if 0 == ret {
		return errors.New("set back main station ipv4 error")
	}
	return nil
}

func getBackMainStationIPv4() (ip string, port uint16) {
	cPort := C.int(0)
	cIp := make([]byte, 16)
	ret := C.getServer(C.int(1), C.int(4), (*C.char)(unsafe.Pointer(&cIp[0])), &cPort)
	_ = ret

	return byteToString(cIp), uint16(cPort)
}

///////////////////////////////////////////
func setMainStationIPv6(ip string, port string) error {
	cs := C.CString(ip)
	pt, _ := strconv.Atoi(port)
	ret := C.setServer(C.int(0), C.int(6), cs, C.int(pt))
	C.free(unsafe.Pointer(cs))
	if 0 == ret {
		return errors.New("set main station ipv6 error")
	}
	return nil
}

func getMainStationIPv6() (ip string, port uint16) {
	cPort := C.int(0)
	cIp := make([]byte, 128)
	ret := C.getServer(C.int(0), C.int(6), (*C.char)(unsafe.Pointer(&cIp[0])), &cPort)
	_ = ret

	return byteToString(cIp), uint16(cPort)
}

func setBackMainStationIPv6(ip string, port string) error {
	cs := C.CString(ip)
	pt, _ := strconv.Atoi(port)
	ret := C.setServer(C.int(1), C.int(6), cs, C.int(pt))
	C.free(unsafe.Pointer(cs))
	if 0 == ret {
		return errors.New("set back main station ipv6 error")
	}
	return nil
}

func getBackMainStationIPv6() (ip string, port uint16) {
	cPort := C.int(0)
	cIp := make([]byte, 128)
	ret := C.getServer(C.int(1), C.int(6), (*C.char)(unsafe.Pointer(&cIp[0])), &cPort)
	_ = ret

	return byteToString(cIp), uint16(cPort)
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

func setSysCPURateThreshold(rate string) error {
	upper, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}
	ret := C.setCpuThreshold(C.int(upper))
	_ = ret
	return nil
}

func getSysCPURateThreshold() string {
	upper := C.int(0)
	ret := C.getCpuThreshold((*C.int)(unsafe.Pointer(&upper)))
	_ = ret
	return strconv.Itoa(int(upper))
}

func setSysMemoryRateThreshold(rate string) error {
	upper, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}
	ret := C.setRamThreshold(C.int(upper))
	_ = ret
	return nil
}

func getSysMemoryRateThreshold() string {
	upper := C.int(0)
	ret := C.getRamThreshold((*C.int)(unsafe.Pointer(&upper)))
	_ = ret
	return strconv.Itoa(int(upper))
}

func setSysDiskRateThreshold(rate string) error {
	upper, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}
	ret := C.setDiskThreshold(C.int(upper))
	_ = ret
	return nil
}

func getSysDiskRateThreshold() string {
	upper := C.int(0)
	ret := C.getDiskThreshold((*C.int)(unsafe.Pointer(&upper)))
	_ = ret
	return strconv.Itoa(int(upper))
}

///////////////////////////////////////////////
func setSysMonitorWndTime(wnd string) error {
	tm, _ := strconv.Atoi(wnd)
	ret := C.setAlarmInterval(C.int(tm))
	_ = ret
	return nil
}

func getSysMonitorWndTime() string {
	tm := C.int(0)
	ret := C.getAlarmInterval(&tm)
	_ = ret
	return strconv.Itoa(int(tm))
}

func setContainerCPURateThreshold(rate string) error {
	rt, _ := strconv.Atoi(rate)
	ret := C.setMonParameter(C.docker, C.cpu, C.int(rt))
	_ = ret
	return nil
}

func getContainerCPURateThreshold() string {
	rt := C.int(0)
	ret := C.getMonParameter(C.docker, C.cpu, &rt)
	_ = ret
	return strconv.Itoa(int(rt))
}

func setContainerMemoryRateThreshold(rate string) error {
	rt, _ := strconv.Atoi(rate)
	ret := C.setMonParameter(C.docker, C.ram, C.int(rt))
	_ = ret
	return nil
}

func getContainerMemoryRateThreshold() string {
	rt := C.int(0)
	ret := C.getMonParameter(C.docker, C.ram, &rt)
	_ = ret
	return strconv.Itoa(int(rt))
}

func setContainerMonitorWndTime(wnd string) error {
	tm, _ := strconv.Atoi(wnd)
	ret := C.setMonParameter(C.docker, C.interval, C.int(tm))
	_ = ret
	return nil
}

func getContainerMonitorWndTime() string {
	tm := C.int(0)
	ret := C.getMonParameter(C.docker, C.interval, &tm)
	_ = ret
	return strconv.Itoa(int(tm))
}

/////////////////////////////////////////////
func setAppCPURateThreshold(rate string) error {
	rt, _ := strconv.Atoi(rate)
	ret := C.setMonParameter(C.app, C.cpu, C.int(rt))
	_ = ret
	return nil
}

func getAppCPURateThreshold() string {
	rt := C.int(0)
	ret := C.getMonParameter(C.app, C.cpu, &rt)
	_ = ret
	return strconv.Itoa(int(rt))
}

func setAppMemoryRateThreshold(rate string) error {
	rt, _ := strconv.Atoi(rate)
	ret := C.setMonParameter(C.app, C.ram, C.int(rt))
	_ = ret
	return nil
}

func getAppMemoryRateThreshold() string {
	rt := C.int(0)
	ret := C.getMonParameter(C.app, C.ram, &rt)
	_ = ret
	return strconv.Itoa(int(rt))
}

func setAppMonitorWndTime(wnd string) error {
	tm, _ := strconv.Atoi(wnd)
	ret := C.setMonParameter(C.app, C.interval, C.int(tm))
	_ = ret
	return nil
}

func getAppMonitorWndTime() string {
	tm := C.int(0)
	ret := C.getMonParameter(C.app, C.ram, &tm)
	_ = ret
	return strconv.Itoa(int(tm))
}

/////////////////////////////////////////////
func setTemperatureThreshold(lower, upper string) error {
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

func getTemperatureThreshold() (lower, upper string) {
	tempUpper := C.int(0)
	tempLower := C.int(0)
	ret := C.getTempThreshold(&tempUpper, &tempLower)
	_ = ret
	return strconv.Itoa(int(tempLower)), strconv.Itoa(int(tempUpper))
}

func setTemperatureThresholdWnd(wnd string) error {
	tm, _ := strconv.Atoi(wnd)
	ret := C.setTempInterval(C.int(tm))
	_ = ret
	return nil
}

func getTemperatureThresholdWnd() string {
	tm := C.int(0)
	ret := C.getTempInterval(&tm)
	_ = ret
	return strconv.Itoa(int(tm))
}

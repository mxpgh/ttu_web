package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	//  "github.com/ziutek/mymysql/mysql"
	//  _ "github.com/ziutek/mymysql/thrsafe"
	"encoding/json"
	_ "log"
)

var (
	adminUser   = "admin"
	adminPasswd = "123456"
)

type Result struct {
	Ret    int
	Reason string
	Data   interface{}
}

type ajaxController struct {
}

func (this *ajaxController) LoginAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := r.ParseForm()
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	userName := strings.TrimSpace(r.FormValue("admin_name"))
	userPassword := r.FormValue("admin_password")

	if userName == "" || userName != adminUser {
		OutputJson(w, 0, "用户名错误", nil)
		return
	}

	if userPassword == "" || userPassword != adminPasswd {
		OutputJson(w, 0, "密码错误", nil)
		return
	}
	/*
	   db := mysql.New("tcp", "", "192.168.100.166", "root", "test", "webdemo")
	   if err := db.Connect(); err != nil {
	       log.Println(err)
	       OutputJson(w, 0, "数据库操作失败", nil)
	       return
	   }
	   defer db.Close()

	   rows, res, err := db.Query("select * from webdemo_admin where admin_name = '%s'", admin_name)
	   if err != nil {
	       log.Println(err)
	       OutputJson(w, 0, "数据库操作失败", nil)
	       return
	   }
	*/

	//name := res.Map("admin_password")
	//admin_password_db := rows[0].Str(name)

	// 存入cookie,使用cookie存储
	cookie := http.Cookie{Name: "admin_name", Value: adminUser, Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)

	OutputJson(w, 1, "操作成功", nil)
	return
}

func (this *ajaxController) ConfigAction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	ipv4 := r.FormValue("IPv4")
	if !validIP(ipv4, "主站IPv4地址错误", w) {
		return
	}
	ipv4Port := r.FormValue("IPv4Port")
	if !validPort(ipv4Port, "主站IPv4端口错误", w) {
		return
	}
	setMainStationIPv4(ipv4, ipv4Port)

	ipv6 := r.FormValue("IPv6")
	if !validIP(ipv6, "主站IPv6地址错误", w) {
		return
	}
	ipv6Port := r.FormValue("IPv6Port")
	if !validPort(ipv6Port, "主站IPv6端口错误", w) {
		return
	}
	setMainStationIPv6(ipv6, ipv6Port)

	backipv4 := r.FormValue("BackIPv4")
	if !validIP(backipv4, "备份主站IPv4地址错误", w) {
		return
	}
	backipv4Port := r.FormValue("BackIPv4Port")
	if !validPort(backipv4Port, "备份主站IPv4端口错误", w) {
		return
	}
	setBackMainStationIPv4(backipv4, backipv4Port)

	backipv6 := r.FormValue("BackIPv6")
	if !validIP(backipv6, "备份主站IPv6地址错误", w) {
		return
	}
	backipv6Port := r.FormValue("BackIPv6Port")
	if !validPort(backipv6Port, "备份主站IPv6端口错误", w) {
		return
	}
	setBackMainStationIPv6(backipv6, backipv6Port)

	sysCPURateUpper := r.FormValue("SysCPURateUpper")
	if !validPercent(sysCPURateUpper, "CPU 使用率上限设置错误", w) {
		return
	}
	setSysCPURateThreshold(sysCPURateUpper)

	sysMemRateUpper := r.FormValue("SysMemRateUpper")
	if !validPercent(sysMemRateUpper, "内存使用上限设置错误", w) {
		return
	}
	setSysMemoryRateThreshold(sysMemRateUpper)

	sysDiskRateUpper := r.FormValue("SysDiskRateUpper")
	if !validPercent(sysDiskRateUpper, "内部存储使用上限设置错误", w) {
		return
	}
	setSysDiskRateThreshold(sysDiskRateUpper)

	sysMonitorWndTime := r.FormValue("SysMonitorWndTime")
	if !validValue(sysMonitorWndTime, "系统监控判断窗口时间设置错误", w) {
		return
	}
	setSysMonitorWndTime(sysMonitorWndTime)

	containerCPURateUpper := r.FormValue("ContainerCPURateUpper")
	if !validPercent(containerCPURateUpper, "容器CPU使用率上限设置错误", w) {
		return
	}
	setContainerCPURateThreshold(containerCPURateUpper)

	containerMemRateUpper := r.FormValue("ContainerMemRateUpper")
	if !validPercent(containerMemRateUpper, "容器内存使用率上限错误", w) {
		return
	}
	setContainerCPURateThreshold(containerMemRateUpper)

	appCPURateUpper := r.FormValue("AppCPURateUpper")
	if !validPercent(appCPURateUpper, "APP CPU使用率上限设置错误", w) {
		return
	}
	setAppCPURateThreshold(appCPURateUpper)

	appMemRateUpper := r.FormValue("AppMemRateUpper")
	if !validPercent(appMemRateUpper, "APP 内存使用率上限设置错误", w) {
		return
	}
	setAppMemoryRateThreshold(appMemRateUpper)

	appMonitorWndTime := r.FormValue("AppMonitorWndTime")
	if !validValue(appMonitorWndTime, "APP 监控判定窗口时间设置错误", w) {
		return
	}
	setAppMonitorWndTime(appCPURateUpper)

	tempLower := r.FormValue("TempLower")
	if !validValue(tempLower, "温度下限阀值错误", w) {
		return
	}
	lowerVal, _ := strconv.Atoi(tempLower)
	if lowerVal < -40 {
		OutputJson(w, 0, "温度下限阀值错误", nil)
		return
	}

	tempUpper := r.FormValue("TempUpper")
	if !validValue(tempUpper, "温度上限阀值设置错误", w) {
		return
	}
	upperVal, _ := strconv.Atoi(tempLower)
	if upperVal > 70 {
		OutputJson(w, 0, "温度上限阀值错误", nil)
		return
	}
	if lowerVal > upperVal {
		OutputJson(w, 0, "温度下限阀值错误", nil)
		return
	}
	setTemperatureThreshold(tempLower, tempUpper)

	tempUpperWnd := r.FormValue("TempUpperWnd")
	if !validValue(tempUpperWnd, "温度上限判定窗口期设置错误", w) {
		return
	}
	setTemperatureThresholdWnd(tempLower)

	log.Println("ConfigAction: ", ipv4)
	OutputJson(w, 0, "修改成功", nil)

	return
}

func (this *ajaxController) UploadAction(w http.ResponseWriter, r *http.Request) {
	//把上传的文件存储在内存和临时文件中
	err := r.ParseMultipartForm(128 << 20)
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	token := r.FormValue("token")
	if token != "" {
		//验证token的合法性
		log.Println("token: ", token)
	} else {
		//不存在token报错
		log.Println("token is nil or empty")
	}

	//获取文件句柄，然后对文件进行存储等处理
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println("form file err: ", err)
		OutputJson(w, 0, "上传错误", nil)
		return
	}

	defer file.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	//创建上传的目的文件
	f, err := os.OpenFile("d:\\tmp\\"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open file err: ", err)
		OutputJson(w, 0, "上传错误", nil)
		return
	}

	defer f.Close()
	//拷贝文件
	io.Copy(f, file)

	OutputJson(w, 0, "上传成功", nil)
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}

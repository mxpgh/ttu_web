package main

import (
	"log"
	"net/http"

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

	admin_name := r.FormValue("admin_name")
	admin_password := r.FormValue("admin_password")

	if admin_name == "" || admin_name != adminUser {
		OutputJson(w, 0, "用户名错误", nil)
		return
	}

	if admin_password == "" || admin_password != adminPasswd {
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

	ipv6 := r.FormValue("IPv6")
	if !validIP(ipv6, "主站IPv6地址错误", w) {
		return
	}
	ipv6Port := r.FormValue("IPv6Port")
	if !validPort(ipv6Port, "主站IPv6端口错误", w) {
		return
	}

	backipv4 := r.FormValue("BackIPv4")
	if !validIP(backipv4, "备份主站IPv4地址错误", w) {
		return
	}
	backipv4Port := r.FormValue("BackIPv4Port")
	if !validPort(backipv4Port, "备份主站IPv4端口错误", w) {
		return
	}

	backipv6 := r.FormValue("BackIPv6")
	if !validIP(backipv6, "备份主站IPv6地址错误", w) {
		return
	}
	backipv6Port := r.FormValue("BackIPv6Port")
	if !validPort(backipv6Port, "备份主站IPv6端口错误", w) {
		return
	}

	sysCPURateUpper := r.FormValue("SysCPURateUpper")
	if !validPercent(sysCPURateUpper, "CPU 使用率上限设置错误", w) {
		return
	}

	sysMemRateUpper := r.FormValue("SysMemRateUpper")
	if !validPercent(sysMemRateUpper, "内存使用上限设置错误", w) {
		return
	}

	sysDiskRateUpper := r.FormValue("SysDiskRateUpper")
	if !validPercent(sysDiskRateUpper, "内部存储使用上限设置错误", w) {
		return
	}

	sysMonitorWndTime := r.FormValue("SysMonitorWndTime")
	if !validValue(sysMonitorWndTime, "系统监控判断窗口时间设置错误", w) {
		return
	}

	containerCPURateUpper := r.FormValue("ContainerCPURateUpper")
	if !validPercent(containerCPURateUpper, "容器CPU使用率上限设置错误", w) {
		return
	}

	containerMemRateUpper := r.FormValue("ContainerMemRateUpper")
	if !validPercent(containerMemRateUpper, "容器内存使用率上限错误", w) {
		return
	}

	appCPURateUpper := r.FormValue("AppCPURateUpper")
	if !validPercent(appCPURateUpper, "APP CPU使用率上限设置错误", w) {
		return
	}

	appMemRateUpper := r.FormValue("AppMemRateUpper")
	if !validPercent(appMemRateUpper, "APP 内存使用率上限设置错误", w) {
		return
	}

	appMonitorWndTime := r.FormValue("AppMonitorWndTime")
	if !validValue(appMonitorWndTime, "APP 监控判定窗口时间设置错误", w) {
		return
	}

	tempUpper := r.FormValue("TempUpper")
	if !validValue(tempUpper, "温度上限阀值设置错误", w) {
		return
	}

	tempUpperWnd := r.FormValue("TempUpperWnd")
	if !validValue(tempUpperWnd, "温度上限判定窗口期设置错误", w) {
		return
	}
	log.Println("ConfigAction: ", ipv4)
	OutputJson(w, 0, "修改成功", nil)

	return
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}

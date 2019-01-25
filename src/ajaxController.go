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
	if !validIP(ipv4) {
		OutputJson(w, 0, "主站IPv4地址错误", nil)
		return
	}
	ipv4Port := r.FormValue("IPv4Port")
	_ = ipv4Port
	if !validPort(ipv4Port) {
		OutputJson(w, 0, "主站IPv4端口错误", nil)
		return
	}
	ipv6 := r.FormValue("IPv6")
	_ = ipv6
	if !validIP(ipv6) {
		OutputJson(w, 0, "主站IPv6地址错误", nil)
		return
	}
	ipv6Port := r.FormValue("IPv6Port")
	_ = ipv6Port
	if !validPort(ipv6Port) {
		OutputJson(w, 0, "主站IPv6端口错误", nil)
		return
	}
	backipv4 := r.FormValue("BackIPv4")
	_ = backipv4
	if !validIP(backipv4) {
		OutputJson(w, 0, "备份主站IPv4地址错误", nil)
		return
	}
	backipv4Port := r.FormValue("BackIPv4Port")
	_ = backipv4Port
	if !validPort(backipv4Port) {
		OutputJson(w, 0, "备份主站IPv4端口错误", nil)
		return
	}
	backipv6 := r.FormValue("BackIPv6")
	_ = backipv6
	if !validIP(backipv6) {
		OutputJson(w, 0, "备份主站IPv6地址错误", nil)
		return
	}
	backipv6Port := r.FormValue("BackIPv6Port")
	_ = backipv6Port
	if !validPort(backipv6Port) {
		OutputJson(w, 0, "备份主站IPv6端口错误", nil)
		return
	}
	sysCPURateUpper := r.FormValue("SysCPURateUpper")
	_ = sysCPURateUpper
	if !validValue(sysCPURateUpper) {
		OutputJson(w, 0, "CPU 使用率上限格式错误", nil)
		return
	}
	sysMemRateUpper := r.FormValue("SysMemRateUpper")
	_ = sysMemRateUpper
	sysDiskRateUpper := r.FormValue("SysDiskRateUpper")
	_ = sysDiskRateUpper
	sysMonitorWndTime := r.FormValue("SysMonitorWndTime")
	_ = sysMonitorWndTime
	containerCPURateUpper := r.FormValue("ContainerCPURateUpper")
	_ = containerCPURateUpper
	containerMemRateUpper := r.FormValue("ContainerMemRateUpper")
	_ = containerMemRateUpper
	appCPURateUpper := r.FormValue("AppCPURateUpper")
	_ = appCPURateUpper
	appMemRateUpper := r.FormValue("AppMemRateUpper")
	_ = appMemRateUpper
	appMonitorWndTime := r.FormValue("AppMonitorWndTime")
	_ = appMonitorWndTime
	tempUpper := r.FormValue("TempUpper")
	_ = tempUpper
	tempUpperWnd := r.FormValue("TempUpperWnd")
	_ = tempUpperWnd
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

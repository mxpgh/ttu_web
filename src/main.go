package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	go timeTask()

	f, err := os.OpenFile(ttuDataFile, os.O_RDONLY, 0600)
	defer f.Close()
	if nil == err {
		passwd := make([]byte, 256)
		l, err := f.Read(passwd)
		if nil == err && l > 0 {
			adminPasswd = string(passwd[:l])
		}
	}

	httpSrv := &http.Server{
		Addr:         ":8888",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.HandleFunc("/basic/", basicHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/status/", statusHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/container/", containerHandler)
	http.HandleFunc("/passwd/", passwdHandler)
	http.HandleFunc("/", NotFoundHandler)
	log.Println("Start ttu_web server: listen port 8888")
	log.Fatal(httpSrv.ListenAndServe())
}

func execBashCmd(bash string) string {
	cmd := exec.Command("/bin/bash", "-c", bash)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func validIP(ip, tip string, w http.ResponseWriter) bool {
	a := net.ParseIP(ip)
	if nil == a {
		OutputJson(w, 0, tip, nil)
		return false
	}
	return true
}

func validPort(port, tip string, w http.ResponseWriter) bool {
	match, err := regexp.MatchString("^[0-9]*$", port)
	if err != nil || false == match {
		OutputJson(w, 0, tip, nil)
		return false
	}

	i, err := strconv.Atoi(port)
	if err != nil {
		OutputJson(w, 0, tip, nil)
		return false
	}
	if i < 0 || i > 65535 {
		OutputJson(w, 0, tip, nil)
		return false
	}

	return true
}

func validValue(val, tip string, w http.ResponseWriter) bool {
	if val != "" {
		_, err := strconv.Atoi(val)
		if err == nil {
			return true
		}

		//if isOK, _ := regexp.MatchString(`^[0-9]\d*|0\.\d*[0-9]\d*$`, val); isOK {
		//	return true
		//}
	}
	OutputJson(w, 0, tip, nil)
	return false
}

func validPercent(val, tip string, w http.ResponseWriter) bool {
	if val != "" {
		if isOK, _ := regexp.MatchString(`^[0-9]\d*|0\.\d*[0-9]\d*$`, val); !isOK {
			OutputJson(w, 0, tip, nil)
			return false
		}
	}
	f, err := strconv.ParseFloat(val, 32)
	if err != nil {
		OutputJson(w, 0, tip, nil)
		return false
	}

	if f < 0.00 || f > 100.00 {
		OutputJson(w, 0, tip, nil)
		return false
	}
	return true
}

func getCurrentTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func fmtTimes(sec int) string {
	days := sec / (24 * 3600)
	hours := sec % (24 * 3600) / 3600
	minutes := sec % 3600 / 60
	seconds := sec % 60
	str := ""
	if days > 0 {
		str += strconv.Itoa(days) + "天"
	}
	if hours > 0 {
		str += strconv.Itoa(hours) + "小时"
	}
	if minutes > 0 {
		str += strconv.Itoa(minutes) + "分"
	}
	if seconds > 0 {
		str += strconv.Itoa(seconds) + "秒"
	}
	return str
}

func genToken() string {
	time := time.Now().Unix()
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time, 10)))
	token := hex.EncodeToString(h.Sum(nil))
	return token
}

func byteToString(c []byte) string {
	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			return string(c[0:i])
		}
	}
	return string(c)
}

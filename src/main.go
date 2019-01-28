package main

import (
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	go timeTask()

	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.HandleFunc("/basic/", basicHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/status/", statusHandler)
	http.HandleFunc("/", NotFoundHandler)
	log.Println("Start ttu_web server: listen port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func execBashCmd(bash string) string {
	cmd := exec.Command("/bin/bash", "-c", bash)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(out)
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
		if isOK, _ := regexp.MatchString(`^[0-9]\d*|0\.\d*[0-9]\d*$`, val); isOK {
			return true
		}
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

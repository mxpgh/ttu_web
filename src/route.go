package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"reflect"
	"strings"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}
	if err != nil {
		log.Println("loginHandler error: ", err)
	} else {
		log.Println("Domain: ", cookie.Domain)
		log.Println("Expires:", cookie.Expires)
		log.Println("Name:", cookie.Name)
		log.Println("Value:", cookie.Value)
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	admin := &adminController{}
	controller := reflect.ValueOf(admin)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ajaxHandler")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}
	log.Println(pathInfo, parts, action)

	ajax := &ajaxController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
	log.Println(action)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	login := &loginController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("basicHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	log.Println(cookie)
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	basic := &basicController{}
	controller := reflect.ValueOf(basic)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("configHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	cfg := &configController{}
	controller := reflect.ValueOf(cfg)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("statusHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	status := &statusController{}
	controller := reflect.ValueOf(status)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("uploadHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	upload := &uploadController{}
	controller := reflect.ValueOf(upload)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func containerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("containerHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	container := &containerController{}
	controller := reflect.ValueOf(container)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("logHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	r.ParseForm()
	page := r.FormValue("page")
	if page == "" {
		page = "0"
	}
	log.Println("page: ", page)
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	log := &logController{}
	controller := reflect.ValueOf(log)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	pageValue := reflect.ValueOf(page)
	method.Call([]reflect.Value{responseValue, requestValue, userValue, pageValue})
}

func passwdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("passwdHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	passwd := &passwdController{}
	controller := reflect.ValueOf(passwd)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func sysmgrHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("sysmgrHandler")
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	sysmgr := &sysmgrController{}
	controller := reflect.ValueOf(sysmgr)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles(path.Join(templatePath, "/html/404.html"))
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

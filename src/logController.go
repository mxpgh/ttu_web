package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type LogItem struct {
	Id       int64
	Date     string
	Filename string
	Function string
	Line     int
	Level    int
	Info     string
	Value    string
}

type Log struct {
	UserName string
	SysTime  string

	Total     int
	Totalpage []int
	Current   int
	LogLists  []LogItem
}

type logController struct {
}

func (this *logController) IndexAction(w http.ResponseWriter, r *http.Request, user, page string) {
	t, err := template.ParseFiles(path.Join(templatePath, "/html/log/index.html"))
	if err != nil {
		log.Println(err)
		return
	}

	var tot int
	var totpage []int
	var logList []LogItem

	db, err := sql.Open("sqlite3", "/root/log/log.db")
	if err != nil {
		log.Println("logControl ", err)
	} else {
		defer db.Close()

		var (
			id       int64
			date     string
			filename string
			function string
			line     int
			level    int
			info     string
			value    string
		)
		err := db.QueryRow("SELECT COUNT(id) FROM log").Scan(&tot)
		if err != nil {
			log.Println("query row: ", err)
		} else {
			log.Println("table log record total: ", tot)
		}
		pg, _ := strconv.Atoi(page)
		totpg := tot / 20
		for i := 1; i < totpg+1; i++ {
			totpage = append(totpage, i)
		}
		rows, err := db.Query("SELECT id, date, filename, function, line, level, info, value FROM log WHERE id > ? LIMIT 20", pg*20)
		log.Println("query: rows->", rows, " err->", err)
		if err != nil {
			log.Println("query: ", err)
		} else {
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&id, &date, &filename, &function, &line, &level, &info, &value)
				log.Println("scan ", err)
				if err != nil {
					log.Println("scan err: ", err)
					break
				} else {
					log.Println(id, "|", date, "|", filename, "|", function, "|", line, "|", info, "|", value)
					item := LogItem{}
					item.Id = id
					item.Date = date
					item.Filename = filename
					item.Function = function
					item.Line = line
					item.Level = level
					item.Info = info
					item.Value = value

					logList = append(logList, item)
				}
			}
			err = rows.Err()
			log.Println("rows: ", err)
			if err != nil {
			}
		}
	}
	/*
			for {
				rows, err := db.Query("SELECT id, date, filename, function, line, level, info, value FROM log WHERE id > ? LIMIT 100", id)
				log.Println("query: rows->", rows, " err->", err)
				if err != nil {
					break
				} else {
					defer rows.Close()
					ret := rows.Next()
					for res := ret; res; res = rows.Next() {
						err = rows.Scan(&id, &date, &filename, &function, &line, &level, &info, &value)
						log.Println("scan ", err)
						if err != nil {
							log.Println("scan err: ", err)
							break
						} else {
							log.Println(id, "|", date, "|", filename, "|", function, "|", line, "|", info, "|", value)
						}
					}
					err = rows.Err()
					log.Println("rows: ", err)
					if err != nil {
						break
					} else {
						if !ret {
							break
						}
					}
				}
			}
		}
	*/
	//log.Println(logList)
	log.Println("log select finish")
	t.Execute(w, &Log{
		user,
		getCurrentTime(),
		tot,
		totpage,
		0,
		logList})
}

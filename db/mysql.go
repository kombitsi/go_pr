package db

import (
	"begeek_bot_golang_12/logger"
	"database/sql"
	"fmt"
	_ "go_pr/loger"
	"go_pr/telegram"
	"time"

)

func mysqlConn(dbName string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root:Passwd@unix(/var/run/mysqld/mysqld.sock)/%s", dbName))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func selectHash(db *sql.DB, hash string) bool {
	var rowHash int
	db.QueryRow("select id from golang_python where link_hash = ?", hash).Scan(&rowHash)
	if rowHash != 0 {
		return true
	}
	return false
}

func insertHash(db *sql.DB, url, page, text, hash string, times int64) {
	resu, err := db.Prepare("INSERT INTO golang_python (site, page_link, page_text, timestamp, link_hash) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		logger.ForError(err)
	}
	_, err = resu.Exec(url, page, text, times, hash)
	if err != nil {
		logger.ForError(err)
	}
}

func CheckSiteNewsBot(url, page, text, hash string) {
	database := mysqlConn("DB_NAME")
	checkLink := selectHash(database, hash)
	times := time.Now().Unix()
	if checkLink == false {
		insertHash(database, url, page, text, hash, times)
		telegram.SendMessage(text)
	}
	defer database.Close()
}

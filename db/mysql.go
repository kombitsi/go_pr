package db

import (
	"begeek_bot/betypes"
	"begeek_bot/logger"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlConn(dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@unix(%s)/%s",
		betypes.BmgeekBotMysqlUser, betypes.BmgeekBotMysqlPasswd, betypes.BmgeekBotMysqlSocker, dbName))
	if err != nil {
		logger.ForError(err)
		return nil, errors.New("не могу подключиться к базе данных")
	}
	return db, nil
}

func InsertUserInfo(chatID, times int, userName, message, command string) {
	database, err := mysqlConn("bmgeek_bot")
	if err != nil {
		logger.ForError(err)
		return
	}
	resu, err := database.Prepare("INSERT INTO user_info (chat_id, username, timestamp, message, command) VALUES (?, ?, ?, ?, ?)")
	logger.ForError(err)
	_, err = resu.Exec(chatID, userName, times, message, command)
	logger.ForError(err)
	defer database.Close()
}

package telegram

import (
	"begeek_bot_golang_12/logger"
	"bytes"
	"fmt"
	_ "go_pr/loger"
	"net/http"
)

const (
	chatID       = "CHAT_ID"
	BOT_TOKEN    = "TOKEN"
	TELEGRAM_URL = "https://api.telegram.org/bot"
)

func SendMessage(text string) {
	textAll := fmt.Sprintf("%s", text)
	data := []byte(fmt.Sprintf(`{"chat_id":%d, "text":"%s", "parse_mode":"HTML", "disable_web_page_preview": true}`, chatID, textAll))
	tx := bytes.NewReader(data)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN), "application/json", tx)
	logger.ForError(err)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"begeek_bot/betypes"
	"begeek_bot/db"
	"begeek_bot/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	webHookInfo := tgbotapi.NewWebhookWithCert(fmt.Sprintf("https://%s:%s/%s", betypes.BOT_ADDRESS, betypes.BOT_PORT, betypes.BOT_TOKEN), betypes.CERT_PATH)
	_, err := bot.SetWebhook(webHookInfo)
	logger.ForError(err)
}

func main() {
	logger.ForError(BotErr)
	setWebhook(NewBot)

	message := func(w http.ResponseWriter, r *http.Request) {
		text, err := ioutil.ReadAll(r.Body)
		logger.ForError(err)
		var botText betypes.BotMessage
		err = json.Unmarshal(text, &botText)
		logger.ForError(err)
		fmt.Println(fmt.Sprintf("%s", text))
		logger.LogFile.Println(fmt.Sprintf("%s", text))

		username := botText.Message.From.Username
		chatUser := botText.Message.From.Id
		chatGroup := botText.Message.Chat.Id
		messageID := botText.Message.Message_id
		mesDate := botText.Message.Date
		userText := botText.Message.Text
		botCommand := strings.Split(botText.Message.Text, "@")[0]
		commandText := strings.Split(botText.Message.Text, " ")

		fmt.Println(username, chatUser, chatGroup, messageID, botCommand, commandText, mesDate)
		db.InsertUserInfo(chatUser, mesDate, username, userText, botCommand)
	}

	http.HandleFunc("/", message)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("%s:%s", betypes.BOT_ADDRESS, betypes.BOT_PORT), betypes.CERT_PATH, betypes.KEY_PATH, nil))
}

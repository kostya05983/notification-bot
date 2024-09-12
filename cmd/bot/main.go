package main

import (
	"log"
	"net/http"
	"os"

	"notification-bot/internal/di"
	session "notification-bot/internal/domain/session"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/dig"
)

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	container := dig.New()
	err := di.InitDi(container)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	container.Invoke(func(sessions session.Repository) {
		updates := bot.ListenForWebhookRespReqFormat(rw, req)

		for update := range updates {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					container.Invoke()

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Для начала работы с ботом нотификаций календаря, нужно авторизоваться в google календаре")

					msg.ReplyToMessageID = update.Message.MessageID

					bot.Send(msg)
				}
			}
		}
	})
}

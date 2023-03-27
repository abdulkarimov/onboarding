package notify

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" //go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
	"log"
)

type AppNotify struct {
	bot *tgbotapi.BotAPI
}

var Notify AppNotify

func New(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	Notify = AppNotify{bot: bot}
}

func (t AppNotify) SendTelegram(email string, url string) {
	msg := tgbotapi.NewMessage(739646468, "Новый пользователь ожидает регистрацию!\n "+email+
		"\n"+url)
	t.bot.Send(msg)
}

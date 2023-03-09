package repositories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" //go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
	"log"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

var Bot TelegramBot

func NewBot(api string) {
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		log.Panic(err)
	}
	Bot = TelegramBot{bot: bot}
}

func (t TelegramBot) SendNotify(email string) {
	t.bot.Send(tgbotapi.NewMessage(1299393358, "Новый пользователь ожидает регистрацию!\n "+email)) //вместо 1299393358 вставить Id человека, которому должно приходить уведомление
}

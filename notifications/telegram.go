package notifications

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendNotifyTelegram(userEmail string, link string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	user_id, e := strconv.ParseInt(os.Getenv("TELEGRAM_USER_ID"), 10, 64)
	if e != nil {
		log.Panic(e)
	}
	if err != nil {
		log.Panic(err)
	}
	bot.Send(tgbotapi.NewMessage(user_id, "👨🏻‍💻Новый пользователь ожидает подтверждение регистрации!\n"+userEmail+"\n"+link))
}

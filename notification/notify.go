package notification

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" //go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
)

type AppNotify struct {
	bot   *tgbotapi.BotAPI
	email smtp.Auth
}

var Notify AppNotify

func New() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_MAIL"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_SERVICE"),
	)

	Notify = AppNotify{bot: bot, email: auth}
}

func (t AppNotify) SendNotifyTelegram(email string, url string) {
	msg := tgbotapi.NewMessage(739646468, "Новый пользователь ожидает регистрацию!\n "+email+
		"\n"+url)
	t.bot.Send(msg)
}

func (t AppNotify) SendNotifyEmail(userMail string, link string) {

	msg := []byte("To: \r\n" +
		"Subject: Новый пользователь\r\n" +
		"\r\n" +
		fmt.Sprintf("Новый пользователь ожидает регистрацию!\n %s \n %s \r\n", userMail, link))

	err := smtp.SendMail(
		os.Getenv("SMTP_SERVICE_PORT"),
		Notify.email,
		os.Getenv("SMTP_MAIL"),
		[]string{"test@gmail.com"},
		msg,
	)

	if err != nil {
		fmt.Println(err)
	}
}

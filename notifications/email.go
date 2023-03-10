package notifications

import (
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
)

func SendNotifyEmail(userEmail string, link string) {
	var e = &email.Email{
		To:      []string{os.Getenv("EMAIL_TO")},
		From:    "Onboarding" + "<" + os.Getenv("EMAL_FROM") + ">",
		Subject: "Новый пользователь Onboarding!",
		Text:    []byte("Новый пользователь ожидает подтверждение регистрации!\n" + userEmail + "\n" + link),
		Headers: textproto.MIMEHeader{},
	}

	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_FROM_PASSWORD"), "smtp.gmail.com"))
}

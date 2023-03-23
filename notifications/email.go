package notifications

import (
	"fmt"
	"net/smtp"
)

func SendNotifyEmail(userMail string, link string) {
	auth := smtp.PlainAuth(
		"",
		"testnotifyonboarding@gmail.com",
		"qigvvmcjlalpcpqi",
		"smtp.gmail.com",
	)

	msg := "Subject: New User\n" + userMail + "\n" + link

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"testnotifyonboarding@gmail.com",
		[]string{"acckarimov@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

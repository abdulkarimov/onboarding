package notifications

import (
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

// const (
// 	smtpAuthAdress   = "smtp.gmail.com"
// 	smtpServerAdress = "smtp.gmai.com:587"
// )

// type SenderEmail interface {
// 	SendEmail(
// 		subject string,
// 		content string,
// 		to []string,
// 		cc []string,
// 		bcc []string,
// 	) error
// }

// type GmailSender struct {
// 	name              string
// 	fromEmailAdress   string
// 	fromEmailPassword string
// }

// func NewGmailSender(name string, fromEmailAdress string, fromEmailPassword string) SenderEmail {
// 	return &GmailSender{
// 		name:              name,
// 		fromEmailAdress:   fromEmailAdress,
// 		fromEmailPassword: fromEmailPassword,
// 	}
// }

// func (sender *GmailSender) SendEmail(
// 	subject string,
// 	content string,
// 	to []string,
// 	cc []string,
// 	bcc []string,

// ) error {
// 	e := email.NewEmail()
// 	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAdress)
// 	e.Subject = subject
// 	e.HTML = []byte(content)
// 	e.To = to
// 	e.Cc = cc
// 	e.Bcc = bcc

// 	smtpAuth := smtp.PlainAuth("", sender.fromEmailAdress, sender.fromEmailPassword, smtpAuthAdress)
// 	return e.Send(smtpServerAdress, smtpAuth)
// }

// func SendNotifyEmail(userMail string, link string) {
// 	sender := NewGmailSender(os.Getenv("EMAIL_FROM_NAME"), os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_FROM_PASSWORD"))

// 	subject := "Новый пользователь ожидает подтверждение регистрации!"
// 	content := userMail + "\n" + link
// 	to := []string{os.Getenv("EMAIL_TO")}
// 	sender.SendEmail(subject, content, to, nil, nil)
// }

func SendNotifyEmail(userMail string, link string) {
	e := email.NewEmail()
	e.From = "Onboarding notfications"
	e.To = []string{os.Getenv("EMAIL_TO")}
	e.Bcc = nil
	e.Cc = nil
	e.Subject = "Новый пользователь ожидает подтверждение регистрации!"
	e.Text = []byte(userMail + "\n" + link)
	e.HTML = nil
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_FROM_PASSWORD"), "smtp.gmail.com"))

}

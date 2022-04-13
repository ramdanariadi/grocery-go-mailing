package config

import (
	_ "embed"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Tunas Dev <>"
const CONFIG_AUTH_EMAIL = ""

//go:embed emailsecret
var CONFIG_AUTH_PASSWORD []byte

func NewMailerAndDialer() (*gomail.Message, *gomail.Dialer) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		string(CONFIG_AUTH_PASSWORD),
	)

	return mailer, dialer
}

func NewMailer() *gomail.Message {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)

	return mailer
}

func NewDialer() *gomail.Dialer {
	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		string(CONFIG_AUTH_PASSWORD),
	)

	return dialer
}

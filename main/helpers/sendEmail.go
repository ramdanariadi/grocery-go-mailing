package helpers

import (
	"github.com/ramdanariadi/grocery-go-mailing/main/config"
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmail(dialer *gomail.Dialer, data MailingData) {
	message := config.NewMailer()
	message.SetHeader("To", data.To...)
	message.SetAddressHeader("Cc", data.Cc.Address, data.Cc.Name)
	message.SetHeader("Subject", data.Subject...)
	message.SetHeader("Body", data.Body)
	err := dialer.DialAndSend(message)
	LogIfError(err)
	log.Println("email sent")
}

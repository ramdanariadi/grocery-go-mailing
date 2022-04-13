package test

import (
	"crypto/tls"
	"fmt"
	"github.com/ramdanariadi/grocery-go-mailing/main/helpers"
	"gopkg.in/gomail.v2"
	"log"
	"testing"
)

func Test_send_email(t *testing.T) {
	mailer := gomail.NewMessage()

	cc := helpers.Cc{
		Address: "",
		Name:    "",
	}

	data := helpers.MailingData{
		To:      []string{""},
		Body:    "",
		Subject: []string{""},
		Cc:      cc,
	}

	from := ""
	username := ""
	password := ""

	mailer.SetHeader("From", fmt.Sprintf("PT. Makmur Subur Jaya <%s>", from))
	mailer.SetHeader("To", data.To...)
	mailer.SetAddressHeader("Cc", data.Cc.Address, data.Cc.Name)
	mailer.SetHeader("Subject", data.Subject...)
	mailer.SetBody("text/html", data.Body)

	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		username,
		password,
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(mailer)
	helpers.LogIfError(err)
	log.Println("Email sent")
}

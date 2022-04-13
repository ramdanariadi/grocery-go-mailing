package main

import (
	"encoding/json"
	"github.com/ramdanariadi/grocery-go-mailing/main/config"
	"github.com/ramdanariadi/grocery-go-mailing/main/helpers"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"log"
	"runtime"
)

func main() {
	const WORKER = 100
	connection, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	helpers.LogIfError(err)

	channel, err := connection.Channel()
	helpers.LogIfError(err)

	queue, err := channel.QueueDeclare(
		"mailing", true, false, false, false, nil)
	helpers.LogIfError(err)

	deliveries, err := channel.Consume(queue.Name, "mail service", true, false, false, false, nil)
	helpers.LogIfError(err)

	forever := make(chan bool)
	messages := make(chan []byte)
	dialer := config.NewDialer()

	go func() {
		for msq := range deliveries {
			//log.Println(msq)
			messages <- msq.Body
		}
	}()
	ctx, cancelFunc := context.WithCancel(context.Background())

	defer func() {
		close(messages)
		close(forever)
		cancelFunc()
	}()
	log.Println(runtime.NumGoroutine())
	for i := 0; i < WORKER; i++ {
		go func(messages <-chan []byte) {
			mailer := config.NewMailer()
			for message := range messages {
				mailingData := helpers.MailingData{}
				err := json.Unmarshal(message, &mailingData)
				helpers.LogIfError(err)
				log.Println(mailingData.Body)
				log.Println(mailingData.To)
				log.Println(mailingData.Cc.Address)
				log.Println(mailingData.Cc.Name)

				mailer.SetHeader("To", mailingData.To...)
				mailer.SetAddressHeader("Cc", mailingData.Cc.Address, mailingData.Cc.Name)
				mailer.SetHeader("Subject", mailingData.Subject...)
				mailer.SetBody("text/html", mailingData.Body)
				err = dialer.DialAndSend(mailer)
				helpers.LogIfError(err)
				mailer.Reset()

				select {
				case <-ctx.Done():
					break
				}
			}
		}(messages)
	}
	log.Println(runtime.NumGoroutine())
	log.Println(runtime.NumCPU())

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

package rabbitmq

import (
	"SendEmail-Service/mail"
	"SendEmail-Service/pkg/config"
	"log"

	"github.com/streadway/amqp"
)

func Consume(queueName string) {
	//rabbitmq url
	amqpServerURL := config.Config("RABBITMQ_URL")

	//  yeni bir rabbitmq bağlantısı oluştur
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()
	// yeni bir kanal oluştur
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// queue deki mesajları dinle
	messages, err := channelRabbitMQ.Consume(
		queueName, // queue name
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully connected to " + queueName + " queue.")

	//channelden gelen mesajlari loopa döndür
	forever := make(chan bool)
	go func() {
		for message := range messages {
			to := []string{string(message.Body)}
			//queue'dan gelen mesaji mail.Send fonksiyonuna gönder
			mail.SendMail(to, queueName)
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-forever
}

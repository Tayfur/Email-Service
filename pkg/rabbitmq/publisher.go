package rabbitmq

import (
	"SendEmail-Service/pkg/config"

	"github.com/streadway/amqp"
)

func Publisher(msg []byte, queueName string) error {
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

	_, err = channelRabbitMQ.QueueDeclare(
		queueName, // queue name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	//yeni mesaj oluştur
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}

	// mesaji publish eder.
	if err := channelRabbitMQ.Publish(
		"",// exchange
		queueName,// queue name
		false,
		false,
		message,
	); err != nil {
		return err
	}

	return nil
}

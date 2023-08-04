package rabbitmq

import (
	"log"
	"rabbitmq-go/helpers"
	"rabbitmq-go/variables"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

func Receive() {
	conn, err := ampq.Dial(variables.RABBITMQ_URL)
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		variables.RABBITMQ_MAIN_QUEUE, // name
		false, // durable
		false, // delete when unusued
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)
	helpers.FailOnError(err, "Failed to declare a queue")

	
	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer
		false, // auto-ack
		false, //exclusive
		false, // no-local
		false, //no-wait
		nil, // args
	)
	helpers.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			time.Sleep(time.Duration(10) * time.Second)
			log.Panicf("Done!")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. to exit press CTRL+C")

	<-forever
}
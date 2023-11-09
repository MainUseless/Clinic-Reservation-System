package inits

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Send(email string, body string){
//open connection to RabbitMQ server
conn, err := amqp.Dial(os.Getenv("rabbitmq_url"))
FailOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

//create channel
ch, err := conn.Channel()
FailOnError(err, "Failed to open a channel")
defer ch.Close()

q, err := ch.QueueDeclare(
	email, // name
	false,   // durable
	false,   // delete when unused
	false,   // exclusive
	false,   // no-wait
	nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err = ch.PublishWithContext(ctx,
	"",     // exchange
	q.Name, // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing {
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
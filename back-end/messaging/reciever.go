package messaging

import (
	"log"
	"os"

	"github.com/gofiber/websocket/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}


func ConsumeAndSendToWebSocket(conn *websocket.Conn , queueName string) {
	conne, err := amqp.Dial(os.Getenv("rabbitmq_url"))
	defer conn.Close()

	ch, err := conne.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
		)
	FailOnError(err, "Failed to declare a queue")

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",                // Consumer
		true,              // Auto-acknowledge messages
		false,             // Exclusive
		false,             // No local
		false,             // No wait
		nil,               // Arguments
	)

	if err != nil {
		log.Println(err)
	}

	// Send messages to WebSocket clients
	for msg := range msgs {
		conn.WriteMessage(websocket.TextMessage, msg.Body)
		log.Printf(" [x] recieved %s\n", msg.Body)
	}

}

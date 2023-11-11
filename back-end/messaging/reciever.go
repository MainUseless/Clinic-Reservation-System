package messaging

import (
	"log"
	"os"

	"github.com/gofiber/websocket/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}


func ConsumeAndSendToWebSocket(conn *websocket.Conn , queueName string) {
	conne, err := amqp.Dial(os.Getenv("rabbitmq_url"))
	defer conn.Close()

	ch, err := conne.Channel()
	defer ch.Close()

	// Consume messages
	msgs, err := ch.Consume(
		queueName, // Queue name
		"",                // Consumer
		true,              // Auto-acknowledge messages
		false,             // Exclusive
		false,             // No local
		false,             // No wait
		nil,               // Arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send messages to WebSocket clients
	for msg := range msgs {
		conn.WriteMessage(websocket.TextMessage, msg.Body)
	}
}

package main

import (
	"log"
	"os"

	"clinic-reservation-system.com/back-end/apis"
	"clinic-reservation-system.com/back-end/inits"
	"clinic-reservation-system.com/back-end/models"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	inits.InitEnv()
	inits.InitDB()

	var User models.User
	if !User.InitTable() {
		log.Fatal("Error in creating users table")
	}

	var Appointment models.Appointment
	if !Appointment.InitTable() {
		log.Fatal("Error in creating appointments table")
	}
}

func consumeAndSendToWebSocket(conn *websocket.Conn) {
	conne, err := amqp.Dial(os.Getenv("rabbitmq_url"))
	defer conn.Close()

	ch, err := conne.Channel()
	defer ch.Close()

	// Consume messages
	msgs, err := ch.Consume(
		"hello", // Queue name
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

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	apis.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// routes := app.GetRoutes()

	// // Print the routes
	// for _, route := range routes {
	// 	fmt.Printf("%s %s\n", route.Method, route.Path)
	// }

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Handle WebSocket connection here
		// Use the RabbitMQ connection to consume messages and send them to the WebSocket connection
		consumeAndSendToWebSocket(c)
	}))

	defer inits.DB.Close()

	app.Listen("127.0.0.1:" + os.Getenv("port"))

}

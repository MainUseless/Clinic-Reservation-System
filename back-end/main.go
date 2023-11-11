package main

import (
	"log"
	"os"

	"clinic-reservation-system.com/back-end/apis"
	"clinic-reservation-system.com/back-end/auth"
	"clinic-reservation-system.com/back-end/inits"
	"clinic-reservation-system.com/back-end/messaging"
	"clinic-reservation-system.com/back-end/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
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

var Email string

func main() {
	var DoctorAuth auth.DoctorAuth
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

	app.Get("/ws",DoctorAuth.Auth ,SetEmail ,websocket.New(func(ctx *websocket.Conn) {
		messaging.ConsumeAndSendToWebSocket(ctx,Email)
	}))

	defer inits.DB.Close()

	app.Listen("127.0.0.1:" + os.Getenv("port"))

}

func SetEmail(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Email = claims["email"].(string)
	return c.Next()
}
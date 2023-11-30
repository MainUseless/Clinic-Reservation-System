package main

import (
	"log"
	"os"
	"strings"

	"clinic-reservation-system.com/back-end/apis"
	// "clinic-reservation-system.com/back-end/auth"
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
// CustomClaims represents the claims structure for JWT
type CustomClaims struct {
	jwt.Claims
	Email string `json:"email"`
}

func main() {
	// var DoctorAuth auth.DoctorAuth
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))
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

	app.Get("/ws",VerifyAndSetEmail,websocket.New(func(ctx *websocket.Conn) {
		messaging.ConsumeAndSendToWebSocket(ctx,Email)
	}))

	defer inits.DB.Close()

	log.Println("secret key is: ", os.Getenv("jwt_secret"))

	app.Listen(":" + os.Getenv("port"))

}

func VerifyAndSetEmail(c *fiber.Ctx) error {
	token,err :=jwt.Parse(c.Query("JWT") , func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwt_secret")), nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		Email = claims["email"].(string)
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Next()
}
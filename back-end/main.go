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
    // "github.com/gofiber/contrib/jwt"
)

func init(){
    inits.InitEnv()
    inits.InitDB()
        
    var User models.User;
    if !User.InitTable(){
        log.Fatal("Error in creating users table")
    }
    
    var Appointment models.Appointment;
    if !Appointment.InitTable(){
        log.Fatal("Error in creating appointments table")
    }
}


func main() {

    app:= fiber.New()
    app.Use(cors.New())
    apis.SetupRoutes(app)


    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    defer inits.DB.Close()

    app.Listen("127.0.0.1:"+os.Getenv("port"))
}

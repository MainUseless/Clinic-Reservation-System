package main

import (
	"os"

	"clinic-reservation-system.com/back-end/apis"
	"clinic-reservation-system.com/back-end/inits"
	"clinic-reservation-system.com/back-end/models"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init(){
    inits.InitEnv()
    inits.InitDB()
        
    var User models.User;
    User.InitTable()
    
    var Appointment models.Appointment;
    Appointment.InitTable()
}


func main() {

    app:= fiber.New()
    app.Use(cors.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Test")
    })
    app.Delete("/", func(c *fiber.Ctx) error {
        inits.DB.Exec("DROP TABLE IF EXISTS appointments,users;")
        return c.SendString("Test")
    })

    defer inits.DB.Close()

    apis.SetupRoutes(app)

    app.Listen(":"+os.Getenv("port"))
}

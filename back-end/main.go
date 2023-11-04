package main

import (
	"fmt"
	"os"

	"clinic-reservation-system.com/back-end/inits"
	"clinic-reservation-system.com/back-end/models"
	"clinic-reservation-system.com/back-end/apis"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init(){
    inits.InitEnv()
    inits.InitDB()

    test:= models.Doctor{Name: "test" , Email: "test", Password: "test"}
    fmt.Print(inits.DB.Create(&test))
}


func main() {

    app:= fiber.New()
    app.Use(cors.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("wsaab")
    })

    apis.SetupRoutes(app)

    app.Listen(":"+os.Getenv("port"))

}

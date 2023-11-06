package apis

import (
	"os"

	"clinic-reservation-system.com/back-end/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/jwt"
)

var accountHandler handlers.AccountHandler
var doctorHandler handlers.DoctorHandler
var patientHandler handlers.PatientHandler

func SetupRoutes(app *fiber.App) {
	api:= app.Group("/api")

	//unauthenticated routes
	signInUp(api)

	//jwt middleware
	api.Use( jwtware.New( jwtware.Config{
        SigningKey: jwtware.SigningKey{Key:[]byte(os.Getenv("jwt_secret"))},
    } ) )

	//authenticated routes
	doctorRoutes(api)
	patientRoutes(api)
}


func signInUp(api fiber.Router){
	api.Route("/account", func(api fiber.Router) {
		api.Post("/signin",accountHandler.SignIn)
		api.Post("/signup",accountHandler.SignUp)
	})

}


func doctorRoutes(api fiber.Router){
	api.Route("/appointment", func(api fiber.Router) {
		api.Post("/", doctorHandler.Auth ,doctorHandler.AddAppointment)
		api.Delete("/", doctorHandler.Auth ,doctorHandler.DeleteAppointment)	
		api.Get("/", doctorHandler.Auth , doctorHandler.GetAppointment)
	})
}


func patientRoutes(api fiber.Router){
	api.Route("/appointment", func(api fiber.Router) {
		api.Post("/", patientHandler.Auth ,patientHandler.ReserveAppointment)
		api.Put("/", patientHandler.Auth ,patientHandler.EditAppointment)
		api.Delete("/", patientHandler.Auth ,patientHandler.DeleteAppointment)	
		api.Get("/", patientHandler.Auth ,patientHandler.GetAppointment)
	})
}
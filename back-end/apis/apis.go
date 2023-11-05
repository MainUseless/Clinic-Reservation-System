package apis


import (
	"clinic-reservation-system.com/back-end/handlers"
	"github.com/gofiber/fiber/v2"
)

var accountHandler handlers.AccountHandler
var doctorHandler handlers.DoctorHandler
var patientHandler handlers.PatientHandler

func SetupRoutes(app *fiber.App) {
	api:= app.Group("/api")

	signInUp(&api)
	doctorRoutes(&api)
	patientRoutes(&api)
}


func signInUp(api *fiber.Router){
	(*api).Post("/signin",accountHandler.SignIn)
	(*api).Post("/signup",accountHandler.SignUp)
}


func doctorRoutes(api *fiber.Router){
	api2 := (*api).Group("/doctor")

	api2.Post("/appointment", doctorHandler.AddAppointment)
	api2.Delete("/appointment", doctorHandler.DeleteAppointment)	
	api2.Get("/appointment", doctorHandler.GetAppointment)
}


func patientRoutes(api *fiber.Router){
	api2 := (*api).Group("/patient")

	api2.Post("/appointment", patientHandler.ReserveAppointment)
	api2.Put("/appointment", patientHandler.EditAppointment)
	api2.Delete("/appointment", patientHandler.DeleteAppointment)	
	api2.Get("/appointment", patientHandler.GetAppointment)

}
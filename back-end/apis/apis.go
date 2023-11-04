package apis


import (
	"clinic-reservation-system.com/back-end/controllers"
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App) {
	signInUp(app)
	doctorRoutes(app)
	patientRoutes(app)
}


func signInUp(app *fiber.App){
	app.Post("/signin",controllers.SignIn)
	app.Post("/signup",controllers.SignUp)
}


func doctorRoutes(app *fiber.App){
	api := app.Group("/doctor")

	api.Post("/appointment", controllers.DoctorAddAppointment)
	api.Delete("/appointment", controllers.DoctorDeleteAppointment)	
	api.Get("/appointment", controllers.DoctorGetAppointment)
}


func patientRoutes(app *fiber.App){
	api := app.Group("/patient")

	api.Post("/appointment", controllers.PatientReserveAppointment)
	api.Put("/appointment", controllers.PatientEditAppointment)
	api.Delete("/appointment", controllers.PatientDeleteAppointment)	
	api.Get("/appointment", controllers.PatientGetAppointment)

}
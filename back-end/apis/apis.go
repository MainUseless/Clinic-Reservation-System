package apis

import (
	"os"

	"clinic-reservation-system.com/back-end/handlers"
	"clinic-reservation-system.com/back-end/auth"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var accountHandler handlers.AccountHandler
var doctorHandler handlers.DoctorHandler
var patientHandler handlers.PatientHandler

var doctorAuth auth.DoctorAuth
var patientAuth auth.PatientAuth

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	//unauthenticated routes
	signInUp(api)

	//jwt middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("jwt_secret"))},
	}))

	//authenticated routes
	doctorRoutes(api)
	patientRoutes(api)
}

func signInUp(api fiber.Router) {
	api.Route("/account", func(api fiber.Router) {
		api.Post("/signin", accountHandler.SignIn)
		api.Post("/signup", accountHandler.SignUp)
	})

}

func doctorRoutes(api fiber.Router) {
	api.Route("/doctor/appointment", func(api fiber.Router) {
		api.Post("/", doctorAuth.Auth, doctorHandler.AddAppointment)
		api.Delete("/", doctorAuth.Auth, doctorHandler.DeleteAppointment)
		api.Get("/", doctorAuth.Auth, doctorHandler.GetAppointment)
	})
}

func patientRoutes(api fiber.Router) {
	api.Get("/patient/doctors", patientAuth.Auth, patientHandler.GetDoctors)
	api.Route("/patient/appointment", func(api fiber.Router) {
		api.Post("/", patientAuth.Auth, patientHandler.ReserveAppointment)
		api.Put("/", patientAuth.Auth, patientHandler.EditAppointment)
		api.Delete("/", patientAuth.Auth, patientHandler.DeleteAppointment)
		api.Get("/", patientAuth.Auth, patientHandler.GetAppointment)
	})
}

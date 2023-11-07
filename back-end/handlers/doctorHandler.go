package handlers

import (
	"time"

	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type DoctorHandler struct{}

func(handler DoctorHandler) Auth(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userType := claims["type"].(string)

	if userType != "doctor" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}


func(handler DoctorHandler) AddAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	id := uint(Tid)
	
	timestamp := ctx.Query("timestamp")

	if timestamp == "" {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map {"error":"Date or time missing"})
	}

	// Parse the date string into a time.Time object
	layout := "2006-01-02 15:04"
	date, err := time.Parse(layout, timestamp)

	if  err != nil || date.Before(time.Now()){
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map {"error":"invalid date or time"})
	}

	appointment  := models.Appointment{ DoctorID: id, AppointmentTime: date }


	if appointment.Create(){
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)

}

func(handler DoctorHandler) GetAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	id := uint(Tid)
    accountType := claims["type"].(string)

	appointment := models.Appointment{ DoctorID:id }

	appointments := appointment.GetAll(accountType)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"appointments": appointments,
	})
}

func(handler DoctorHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorDeleteAppointment")
}



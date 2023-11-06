package handlers

import (
	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type PatientHandler struct{}


func(handler PatientHandler) Auth(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userType := claims["type"].(string)

	if userType != "patient" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}



func(handler PatientHandler) ReserveAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientReserveAppointment")
}

func(handler PatientHandler) GetAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
    id := claims["id"].(uint)
    accountType := claims["type"].(string)

	appointment := models.Appointment{ PatientID:id }

	appointments := appointment.GetAll(accountType)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"appointments": appointments,
	})

}

func(handler PatientHandler) EditAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientEditAppointment")
}

func(handler PatientHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientDeleteAppointment")
} 
package handlers

import (
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
	return ctx.SendString("PatientGetAppointment")
}

func(handler PatientHandler) EditAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientEditAppointment")
}

func(handler PatientHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientDeleteAppointment")
} 
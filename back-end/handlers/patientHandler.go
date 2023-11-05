package handlers

import "github.com/gofiber/fiber/v2"

type PatientHandler struct{}

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
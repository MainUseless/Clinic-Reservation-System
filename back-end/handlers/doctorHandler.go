package handlers

import "github.com/gofiber/fiber/v2"

type DoctorHandler struct{}

func(handler DoctorHandler) AddAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorAddAppointment")
}

func(handler DoctorHandler) GetAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorGetAppointment")
}

func(handler DoctorHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorDeleteAppointment")
} 
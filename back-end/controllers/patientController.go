package controllers

import "github.com/gofiber/fiber/v2"

func PatientReserveAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientReserveAppointment")
}

func PatientGetAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientGetAppointment")
}

func PatientEditAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientEditAppointment")
}

func PatientDeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("PatientDeleteAppointment")
} 
package controllers

import "github.com/gofiber/fiber/v2"

func DoctorAddAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorAddAppointment")
}

func DoctorGetAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorGetAppointment")
}

func DoctorDeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorDeleteAppointment")
} 
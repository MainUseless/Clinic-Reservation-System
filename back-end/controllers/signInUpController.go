package controllers

import "github.com/gofiber/fiber/v2"

func SignIn(ctx *fiber.Ctx) error {
	return ctx.SendString("SignIn")
}

func SignUp(ctx *fiber.Ctx) error {
	return ctx.SendString("SignUp")
}
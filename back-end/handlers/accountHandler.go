package handlers

import "github.com/gofiber/fiber/v2"

type AccountHandler struct{}

func(handler AccountHandler) SignIn(ctx *fiber.Ctx) error {
	return ctx.SendString("SignIn")
}

func(handler AccountHandler) SignUp(ctx *fiber.Ctx) error {
	return ctx.SendString("SignUp")
}
package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type DoctorAuth struct{}

func (doctorAuth DoctorAuth) Auth(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userType := claims["type"].(string)

	if userType != "doctor" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
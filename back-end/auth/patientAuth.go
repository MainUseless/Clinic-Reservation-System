package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type PatientAuth struct{}

func (patientAuth PatientAuth) Auth(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userType := claims["type"].(string)

	if userType != "patient" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}

package handlers

import (
	"os"

	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AccountHandler struct{}

func(handler AccountHandler) SignIn(ctx *fiber.Ctx) error {
	var account models.User
	email := ctx.Query("email")
	password := ctx.Query("password")

	if email == "" || password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Email or password missing"})
	}

	account.Email = email
	account.Password = password

	if !account.Get(){
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Account not found"})
	}

	claims := jwt.MapClaims{
		"id":    account.ID,
		"name":  account.Name,
		"type":  account.Type,
		"email": account.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Error in signing token"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": signedToken,
	})

}

func(handler AccountHandler) SignUp(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	password := ctx.Query("password")
	userType := ctx.Query("type")
	name := ctx.Query("name")

	if email == "" || password == "" || userType == "" || name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Missing fields"})
	}

	var account models.User 
	
	account.Email = email
	account.Password = password
	account.Type = userType
	account.Name = name

	if account.Create(){
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Account created successfully",
		})	
	}else{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Account already exists or error in creating account",
		})
	}
	
}
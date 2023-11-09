package handlers

import (
	"log"
	"os"

	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AccountHandler struct{}

func(handler AccountHandler) SignIn(ctx *fiber.Ctx) error {
	user := new(models.User)

    if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Missing fields",
		})
    }

	if user.Email == "" || user.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Missing fields"})
	}

	password := user.Password

	if !user.Get(){
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Account not found"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Incorrect email or password"})
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"type":  user.Type,
		"email": user.Email,
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
	user := new(models.User)

    if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Missing fields",
		})
    }

	if user.Email == "" || user.Password == "" || user.Type == "" || user.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Missing fields"})
	}
		
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error in hashing password",
		})
	}
	log.Println(string(hashedPassword))
	user.Password = string(hashedPassword)

	if user.Create(){
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Account created successfully",
		})	
	}else{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Account already exists or error in creating account",
		})
	}
	
}
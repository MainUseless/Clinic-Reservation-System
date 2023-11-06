package handlers

import (
	"time"

	"clinic-reservation-system.com/back-end/inits"
	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type DoctorHandler struct{}


func(handler DoctorHandler) Auth(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userType := claims["type"].(string)

	if userType != "doctor" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}


func(handler DoctorHandler) AddAppointment(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(uint)
	
	reqDate := ctx.Query("date")
	reqTime := ctx.Query("time")

	if reqDate == "" || reqTime == "" {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map {"error":"Date or time missing"})
	}

	date, err := time.Parse("2006-01-02 15:03", reqDate+" "+reqTime)

	if  err != nil || date.Before(time.Now()){
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map {"error":"invalid date or time"})
	}

	query := `
	SELECT EXISTS (
		SELECT 1
		FROM appointments
		WHERE doctor_id=? and ABS(TIMESTAMPDIFF(HOUR, timestamp_column, ?)) != 1
	) AS result;
	`
	
	var result bool

	err = inits.DB.QueryRow(query, id, date).Scan(&result)

	if err != nil || !result {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map {"error":"Error in checking appointment"})
	}

	appointment  := models.Appointment{ DoctorID: id, AppointmentTime: date }

	if appointment.Create(){
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": result,
		})
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)

}

func(handler DoctorHandler) GetAppointment(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)


	return ctx.SendString("DoctorGetAppointment "+name)
}

func(handler DoctorHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	return ctx.SendString("DoctorDeleteAppointment")
}



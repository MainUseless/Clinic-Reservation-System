package handlers

import (
	"database/sql"
	"strconv"

	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type PatientHandler struct{}

func (handler PatientHandler) ReserveAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	id := uint(Tid)
	nullableID := sql.NullInt64{Int64: int64(id), Valid: true}

	appointment_id := ctx.Query("appointment_id")

	intAppointmentID, _ := strconv.Atoi(appointment_id)
	appointment := models.Appointment{PatientID: nullableID, ID: sql.NullInt64{Int64: int64(intAppointmentID), Valid: true}}

	if appointment.Reserve() {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)

}

func (handler PatientHandler) GetAppointment(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	id := uint(Tid)
	nullableID := sql.NullInt64{Int64: int64(id), Valid: true}

	doctorID := ctx.Query("doctor_id")

	appointment := models.Appointment{PatientID: nullableID}

	var appointments []models.Appointment

	if doctorID == "" {
		appointments = appointment.GetAll("patient", true)
	} else {
		intDoctorID, _ := strconv.Atoi(doctorID)
		nullableDoctorID := sql.NullInt64{Int64: int64(intDoctorID), Valid: true}
		appointment.DoctorID = nullableDoctorID
		appointments = appointment.GetAll("doctor", true)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"appointments": appointments,
	})

}

func (handler PatientHandler) EditAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	id := uint(Tid)
	nullableID := sql.NullInt64{Int64: int64(id), Valid: true}

	appointmentID, err := strconv.Atoi(ctx.Query("appointment_id"))
	timeStamp := ctx.Query("timestamp")

	if err != nil || timeStamp == "" {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "Required fields missing"})
	}

	nullableAppointmentID := sql.NullInt64{Int64: int64(appointmentID), Valid: true}
	appointment := models.Appointment{PatientID: nullableID, ID: nullableAppointmentID, AppointmentTime: sql.NullString{String: timeStamp, Valid: true}}

	if appointment.Edit() {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)
}

func (handler PatientHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	id := uint(Tid)
	nullableID := sql.NullInt64{Int64: int64(id), Valid: true}

	appointmentID, err := strconv.Atoi(ctx.Query("appointment_id"))

	if err != nil {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "Appointment ID missing"})
	}

	nullableAppointmentID := sql.NullInt64{Int64: int64(appointmentID), Valid: true}
	appointment := models.Appointment{PatientID: nullableID, ID: nullableAppointmentID}

	if appointment.UnReserve() {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)
}

package handlers

import (
	"database/sql"
	"strconv"

	"clinic-reservation-system.com/back-end/messaging"
	"clinic-reservation-system.com/back-end/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type PatientHandler struct{}

func (handler PatientHandler) GetDoctors(ctx *fiber.Ctx) error {
	doctor := models.User{Type: "doctor"}

	doctors := doctor.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"doctors": doctors,
	})
}

func (handler PatientHandler) ReserveAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	nullableID := sql.NullInt64{Int64: int64(Tid), Valid: true}

	appointment_id := ctx.Query("appointment_id")

	intAppointmentID, _ := strconv.Atoi(appointment_id)
	appointment := models.Appointment{PatientID: nullableID, ID: sql.NullInt64{Int64: int64(intAppointmentID), Valid: true}}

	if appointment.Reserve() {
		//sender send email to doctor
		doctorEmail := appointment.GetDoctorContact()
		messaging.Send(doctorEmail, `{doctor_id:`+strconv.Itoa(int(appointment.DoctorID.Int64))+`,patient_id:`+strconv.Itoa(int(Tid))+`,operation:"ReservationCreated"}`)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusBadRequest)

}

func (handler PatientHandler) GetAppointment(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	nullableID := sql.NullInt64{Int64: int64(Tid), Valid: true}

	doctorID := ctx.Query("doctor_id")

	appointment := models.Appointment{PatientID: nullableID}

	var appointments []fiber.Map

	if doctorID == "" {
		appointments = appointment.GetReserved("patient")
	} else {
		intDoctorID, _ := strconv.Atoi(doctorID)
		appointment.DoctorID = sql.NullInt64{Int64: int64(intDoctorID), Valid: true}
		appointments = appointment.GetAll("doctor")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"appointments": appointments,
	})

}

func (handler PatientHandler) EditAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	nullableID := sql.NullInt64{Int64: int64(Tid), Valid: true}

	appointmentID, err := strconv.Atoi(ctx.Query("appointment_id"))
	new_appointmentID, err1 := strconv.Atoi(ctx.Query("new_appointment_id"))

	if err != nil {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "Appointment ID missing"})
	}

	if err1 != nil {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "New Appointment ID missing"})
	}

	nullableAppointmentID := sql.NullInt64{Int64: int64(appointmentID), Valid: true}
	appointment := models.Appointment{PatientID: nullableID, ID: nullableAppointmentID}

	doctorEmail := appointment.GetDoctorContact()

	if doctorEmail == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unable to get doctor contact"})
	}

	if !appointment.UnReserve(){
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unable to unreserve appointment"})
	}

	appointment.ID = sql.NullInt64{Int64: int64(new_appointmentID), Valid: true}

	if appointment.Reserve() {
		//sender send email to doctor
		messaging.Send(doctorEmail, `{doctor_id:`+strconv.Itoa(int(appointment.DoctorID.Int64))+`,patient_id:`+strconv.Itoa(int(Tid))+`,operation:"ReservationEdited"}`)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusBadRequest)
}

func (handler PatientHandler) DeleteAppointment(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	Tid := claims["id"].(float64)
	nullableID := sql.NullInt64{Int64: int64(Tid), Valid: true}

	appointmentID, err := strconv.Atoi(ctx.Query("appointment_id"))

	if err != nil {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "Appointment ID missing"})
	}

	nullableAppointmentID := sql.NullInt64{Int64: int64(appointmentID), Valid: true}
	appointment := models.Appointment{PatientID: nullableID, ID: nullableAppointmentID}

	if appointment.UnReserve() {
		//sender send email to doctor
		doctorEmail := appointment.GetDoctorContact()
		messaging.Send(doctorEmail, `{doctor_id:`+strconv.Itoa(int(appointment.DoctorID.Int64))+`,patient_id:`+strconv.Itoa(int(Tid))+`,operation:"ReservationCancelled"}`)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": true,
		})
	}

	return ctx.SendStatus(fiber.StatusBadRequest)
}

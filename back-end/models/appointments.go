package models

import (
	"time"
	"clinic-reservation-system.com/back-end/inits"
)

type Appointment struct {
	DoctorID        uint
	PatientID       uint
	AppointmentTime time.Time
}

func (a Appointment) InitTable() bool {
	query := `
	CREATE TABLE IF NOT EXISTS appointments(
		id SERIAL PRIMARY KEY,
		doctor_id int REFERENCES users(id),
		patient_id int REFERENCES users(id),
		appointment_time timestamp
		);
		`
	_, err := inits.DB.Exec(query)

	return err == nil

}

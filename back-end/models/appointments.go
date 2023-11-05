package models

import (
	"log"
	"time"

	"clinic-reservation-system.com/back-end/inits"
)

type Appointment struct {
	DoctorID        uint `json:"doctor_id"`
	PatientID       uint `json:"patient_id"`
	AppointmentTime time.Time `json:"appointment_time"`
}

func (a Appointment) InitTable() bool {
	query := `
	CREATE TABLE IF NOT EXISTS appointments(
		id int NOT NULL AUTO_INCREMENT,
		doctor_id int NOT NULL,
		patient_id int,
		appointment_time timestamp	NOT NULL,
		PRIMARY KEY (id),
		KEY doctor_id_fk (doctor_id),
		KEY patient_id_fk (patient_id)
		);
		`

	_, err := inits.DB.Exec(query)

	log.Println(err)

	return err == nil

}

func (a Appointment) Create() bool {
	query := `
	INSERT INTO appointments(doctor_id,appointment_time) VALUES($1,$2);
	`
	_, err := inits.DB.Exec(query, a.DoctorID, a.AppointmentTime)

	return err == nil
}


func (a Appointment) Delete() bool {
	query := `
	DELETE FROM appointments WHERE doctor_id=$1 AND appointment_time=$2;
	`
	_, err := inits.DB.Exec(query, a.DoctorID, a.AppointmentTime)

	return err == nil
}

func (a Appointment) GetAll(userType string) []Appointment{
	var query string

	if userType == "doctor" {
		query = `
		SELECT * FROM appointments WHERE doctor_id=$1;
		`
	}else{
		query = `
		SELECT * FROM appointments WHERE patient_id=$1;
		`
	}


	rows, err := inits.DB.Query(query, a.DoctorID, a.AppointmentTime)
	var appointments []Appointment
	
	if err != nil {
		log.Fatal(err)
		return appointments
	}
	
	defer rows.Close()


	for rows.Next() {
		var appointment Appointment
		err = rows.Scan(&appointment.DoctorID, &appointment.PatientID, &appointment.AppointmentTime)
		if err != nil {
			log.Println(err)
			return appointments
		}
		appointments = append(appointments, appointment)
	}


	return appointments

}
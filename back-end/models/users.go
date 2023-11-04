package models

import (
	"clinic-reservation-system.com/back-end/inits"
)

type User struct {
	Name string 
	Type string	

	//should be moved to a separate accounts model to be hashed
	Email string
	Password string

	Appointments []Appointment
}

func( u User ) InitTable() bool {
	query:= `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name varchar(30) NOT NULL,
		type ENUM('doctor','patient') NOT NULL,
		email varchar(30),
		password varchar(30),
		appointments_fk int REFERENCES appointments(id)
		);
		`
	_,err := inits.DB.Exec(query)
	
	return err == nil

}
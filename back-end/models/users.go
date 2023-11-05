package models

import (
	"log"

	"clinic-reservation-system.com/back-end/inits"
)

type User struct {
	Name string		`json:"name"`
	Type string		`json:"type"`

	//should be moved to a separate accounts model to be hashed
	Email string	`json:"email"`
	Password string	`json:"password"`

	Appointments []Appointment `json:"appointments"`
}

func( u User ) InitTable() bool {
	query:= `
	CREATE TABLE IF NOT EXISTS users(
		id int NOT NULL AUTO_INCREMENT,
		name varchar(30) NOT NULL,
		type ENUM('doctor','patient') NOT NULL,
		email varchar(30) NOT NULL UNIQUE,
		password varchar(30) NOT NULL,
		PRIMARY KEY (id)
		);
		`
	_,err := inits.DB.Exec(query)
	
	log.Println(err)

	return err == nil

}

func (u User) Create() bool {
	query := `
	INSERT INTO users(name,type,email,password) VALUES($1,$2,$3,$4);
	`
	_,err := inits.DB.Exec(query,u.Name,u.Type,u.Email,u.Password)

	return err == nil
}

func (u User) Get() User {
	query := `
	SELECT * FROM users WHERE email=$1 AND password=$2;
	`
	row := inits.DB.QueryRow(query,u.Email,u.Password)

	var user User

	row.Scan(&user.Name,&user.Type,&user.Email,&user.Password)

	return user
}
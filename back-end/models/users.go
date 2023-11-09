package models

import (
	"log"

	"clinic-reservation-system.com/back-end/inits"
)

type User struct {
	ID uint			`json:"id"`
	Name string		`json:"name"`
	Type string		`json:"type"`

	//should be moved to a separate accounts model to be hashed
	Email string	`json:"email"`
	Password string	`json:"password"`
}

func( u User ) InitTable() bool {
	query:= `
	CREATE TABLE IF NOT EXISTS users(
		id int NOT NULL AUTO_INCREMENT,
		name varchar(30) NOT NULL,
		type ENUM('doctor','patient') NOT NULL,
		email varchar(30) NOT NULL UNIQUE,
		password varchar(72) NOT NULL,
		PRIMARY KEY (id)
		);
		`
	_,err := inits.DB.Exec(query)

	return err == nil

}

func (u *User) Create() bool {
	query := `
	INSERT INTO users(name,type,email,password) VALUES(?,?,?,?);
	`
	row,err := inits.DB.Exec(query,u.Name,u.Type,u.Email,u.Password)
	
	if  err!= nil {
		log.Println("Error in creating user in database")
		log.Println(err.Error())
		return false
	}
	
	var Tid int64
	Tid,err = row.LastInsertId()
	
	id := uint(Tid)

	if err != nil {
		log.Println("Error in getting last inserted id")
		log.Println(err.Error())
		return false
	}else{
		(*u).ID = id
		return true
	}

}

func (u *User) Get() bool {
	query := `
	SELECT * FROM users WHERE email=?;
	`

	err := inits.DB.QueryRow(query,u.Email).Scan(&u.ID,&u.Name,&u.Type,&u.Email,&u.Password)

	if err != nil {
		log.Println("Error in getting user from database")
		log.Println(err.Error())
		return false
	}else{
		return true
	}

}


func (u *User) GetName(id uint) bool {
	query := `
	SELECT name FROM users WHERE id=?;
	`

	err := inits.DB.QueryRow(query,&u.ID).Scan(&u.Name)

	if err != nil {
		log.Println("Error in getting user from database")
		log.Println(err.Error())
		return false
	}else{
		return true
	}

}
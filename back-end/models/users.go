package models

import (
	"log"

	"clinic-reservation-system.com/back-end/inits"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	//should be moved to a separate accounts model to be hashed
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) InitTable() bool {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id int NOT NULL AUTO_INCREMENT,
		name varchar(30) NOT NULL,
		type ENUM('doctor','patient') NOT NULL,
		email varchar(30) NOT NULL UNIQUE,
		password varchar(72) NOT NULL,
		PRIMARY KEY (id)
		);
		`
	_, err := inits.DB.Exec(query)

	return err == nil

}

func (u User) GetAll() []User {
	query := `
	SELECT id,name FROM users WHERE type=?;
	`

	rows, err := inits.DB.Query(query, u.Type)

	if err != nil {
		log.Println("Error in getting users from database")
		log.Println(err.Error())
		return nil
	}

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)

		if err != nil {
			log.Println("Error in getting users from database")
			log.Println(err.Error())
			return nil
		}

		users = append(users, user)
	}

	return users
}

func (u *User) Create() bool {
	query := `
	INSERT INTO users(name,type,email,password) VALUES(?,?,?,?);
	`
	row, err := inits.DB.Exec(query, u.Name, u.Type, u.Email, u.Password)

	if err != nil {
		log.Println("Error in creating user in database")
		log.Println(err.Error())
		return false
	}

	var Tid int64
	Tid, err = row.LastInsertId()

	id := uint(Tid)

	if err != nil {
		log.Println("Error in getting last inserted id")
		log.Println(err.Error())
		return false
	} else {
		(*u).ID = id
		return true
	}

}

func (u *User) Get() bool {
	query := `
	SELECT * FROM users WHERE email=? or id=?;
	`

	err := inits.DB.QueryRow(query, u.Email , u.ID).Scan(&u.ID, &u.Name, &u.Type, &u.Email, &u.Password)

	if err != nil {
		log.Println("Error in getting user from database")
		log.Println(err.Error())
		return false
	} else {
		return true
	}

}

func GetName(id uint) string {
	query := `
	SELECT name FROM users WHERE id=?;
	`
	var name string;
	err := inits.DB.QueryRow(query, id).Scan(&name)

	if err != nil {
		log.Println("Error in getting user from database")
		log.Println(err.Error())
		return ""
	} else {
		return name
	}

}

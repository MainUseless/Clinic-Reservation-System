package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name string `gorm:"not null"`

	//should be moved to a separate accounts model to be hashed
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`

	Appointments []Appointment	`gorm:"foreignkey:DoctorID"`
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorID   uint      `gorm:"not null"`
	Doctor Doctor 		
	PatientID  uint      `gorm:"not null"`
	Patient Patient 	
	AppointmentTime time.Time `gorm:"not null"`
}
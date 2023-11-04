package inits

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"clinic-reservation-system.com/back-end/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}	
	
	DB.Migrator().DropTable("doctors")
	DB.Migrator().DropTable("patients")
	DB.Migrator().DropTable("appointments")

	DB.AutoMigrate(&models.Doctor{},&models.Patient{},&models.Appointment{})

	fmt.Println("Migrated");
}
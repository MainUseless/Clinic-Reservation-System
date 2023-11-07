package inits

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}

}

func reset(){
	DB.Query("drop table if exists appointments,users,doctors,patients;")
}

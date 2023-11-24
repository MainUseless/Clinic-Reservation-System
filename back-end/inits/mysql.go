package inits

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	if(os.Getenv("mysql_url")==""){
		log.Fatal("mysql_url is not set")
	}
	var err error
	DB, err = sql.Open("mysql", os.Getenv("mysql_url"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
}

func reset() {
	DB.Query("drop table if exists appointments,users,doctors,patients;")
}

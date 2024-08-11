package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Bank_connection_string = ""
	Port                   = 0
	Port_db                = 0
	Host_db                = ""
	User_db                = ""
	Password_db            = ""
	Database               = ""
)

func InitializeEnv() {
	var err error

	if err = godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 8080
	}

	Port_db, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	if err != nil {
		Port = 5432
	}

	Host_db = os.Getenv("POSTGRES_HOST")
	User_db = os.Getenv("POSTGRES_USER")
	Password_db = os.Getenv("POSTGRES_PASSWORD")
	Database = os.Getenv("POSTGRES_DB")

	Bank_connection_string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host_db, Port_db, User_db, Password_db, Database)
}

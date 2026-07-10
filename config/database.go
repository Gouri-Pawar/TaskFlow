package config

import (
		"fmt"
		"log"
		"os"

		"github.com/joho/godotenv"
		"gorm.io/driver/postgres"
		"gorm.io/gorm"
)

var DB *gorm.DB			 //global variable that stores our DB connection

func ConnectDB() {

err := godotenv.Load() 		//reads data from .env file
if err != nil {
	log.Fatal("Error loading .env file")
}

dsn := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
)

// gorm.Open() --> Go appn --> gorm --> postgresql driver and then Postgresql DB if connection succeeds database contains connection object

database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatal(err)
}

DB = database   		//storing connection

fmt.Println("Database Connected Successfully")

}
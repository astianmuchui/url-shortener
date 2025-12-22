package db

import (
	"fmt"
	"github.com/astianmuchui/url-shortener/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB

func init() {
	env.Load()

	Connect()
}

func Connect() {

	var err error

	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Println("Error parsing environment variables")
	}

	tz := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", dbhost, dbuser, dbpassword, dbname, port, tz)

	DB, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: false,
		}), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	log.Println("Connected to database successfully")

}

package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"bookstore_db",
		5432,
	)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db = d
	fmt.Println("Database connection established")

}

func GetDb() *gorm.DB {
	return db
}

package pkg

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	connectionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func NewPostgresConection() (db *gorm.DB) {
	dsn := fmt.Sprintf(connectionString, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

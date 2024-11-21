package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// PostgreSQL connection string: "host=<HOST> user=<USER> password=<PASSWORD> dbname=<DBNAME> port=<PORT> sslmode=<SSLMODE>"
	dsn := "host=localhost user=postgres password=gozayaan dbname=book-store port=5433 sslmode=disable TimeZone=Asia/Dhaka"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

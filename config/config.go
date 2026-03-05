package config

import (
	"challengeGO/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func InitDB() *gorm.DB {

	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "johannes")
	dbPass := getEnv("DB_PASSWORD", "mypassword123")
	dbName := getEnv("DB_NAME", "challengego")
	dbPort := getEnv("DB_PORT", "5432")
	dbSSL := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Book{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected successfully")
	return db
}

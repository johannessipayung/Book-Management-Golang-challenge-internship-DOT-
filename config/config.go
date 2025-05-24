package config

import (
	"challengeGO/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=johannes password=mypassword123 dbname=challengego port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Category{}, // 👉 Migrasi Category dulu sebelum Book
		&model.Book{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

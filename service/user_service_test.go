package service

import (
	"challengeGO/model"
	"challengeGO/repository"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupService(t *testing.T) UserService {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed open db: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("failed migrate: %v", err)
	}

	repo := repository.NewUserRepository(db)
	return NewUserService(repo)
}

func TestRegisterUser(t *testing.T) {
	service := setupService(t)

	user := model.User{
		Username: "serviceTest",
		Email:    "service@mail.com",
		Password: "123",
	}

	err := service.Register(&user)
	if err != nil {
		t.Fatalf("failed register: %v", err)
	}
}

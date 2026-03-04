package repository

import (
	"challengeGO/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed open db: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("failed migrate: %v", err)
	}

	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	user := model.User{
		Username: "repoTest",
		Email:    "repo@mail.com",
		Password: "123",
	}

	err := repo.Create(&user)
	if err != nil {
		t.Fatalf("failed create user: %v", err)
	}
}

func TestGetByEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	user := model.User{
		Username: "repoTest2",
		Email:    "repo2@mail.com",
		Password: "123",
	}

	_ = repo.Create(&user)

	found, err := repo.FindByEmail("repo2@mail.com")
	if err != nil {
		t.Fatalf("failed get user: %v", err)
	}

	if found.Email != user.Email {
		t.Fatalf("email mismatch")
	}
}

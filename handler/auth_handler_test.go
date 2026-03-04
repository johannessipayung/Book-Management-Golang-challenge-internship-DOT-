package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"challengeGO/model"
	"challengeGO/repository"
	"challengeGO/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupHandler(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed db: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("failed migrate: %v", err)
	}

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := NewAuthHandler(service)

	r := gin.Default()
	r.POST("/register", handler.Register)

	return r
}

func TestRegisterHandler(t *testing.T) {
	r := setupHandler(t)

	user := model.User{
		Username: "handlerTest",
		Email:    "handler@mail.com",
		Password: "123",
	}

	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", w.Code)
	}
}

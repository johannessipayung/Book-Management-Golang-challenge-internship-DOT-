package test

import (
	"bytes"
	"challengeGO/handler"
	"challengeGO/model"
	"challengeGO/repository"
	"challengeGO/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupAuthTest(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	r := gin.Default()
	authHandler := handler.NewAuthHandler(
		service.NewUserService(
			repository.NewUserRepository(db),
		),
	)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	return r, db
}

func TestRegisterLogin(t *testing.T) {
	r, _ := setupAuthTest(t)

	user := model.User{
		Username: "TestAuth",
		Email:    "testauth@mail.com",
		Password: "pass",
	}

	body, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("failed to marshal user: %v", err)
	}

	// ================= REGISTER =================
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// ================= LOGIN =================
	loginReq := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	loginReq.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	assert.NotEmpty(t, resp["token"])
}

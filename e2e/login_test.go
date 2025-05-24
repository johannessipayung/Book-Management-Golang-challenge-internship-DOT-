package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"challengeGO/config"
	"challengeGO/handler"
	"challengeGO/repository"
	"challengeGO/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Fungsi untuk hapus user dengan username/email tertentu sebelum test agar tidak duplicate
func cleanupTestUser(username, email string) {
	db := config.InitDB()
	userRepo := repository.NewUserRepository(db)
	// Gunakan repo langsung delete user dengan username atau email yang sama
	userRepo.DeleteByUsernameOrEmail(username, email)
}

func setupRouter() *gin.Engine {
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	r := gin.Default()

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	return r
}

func TestLoginSuccess(t *testing.T) {
	// Generate username/email unik pakai timestamp supaya gak duplikat
	username := fmt.Sprintf("testuser_%d", time.Now().UnixNano())
	email := fmt.Sprintf("test_%d@example.com", time.Now().UnixNano())
	password := "testpass"

	// Bersihkan dulu user lama dengan username/email yang sama, jika ada
	cleanupTestUser(username, email)

	r := setupRouter()

	// Register user terlebih dahulu
	registerBody := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}
	registerJSON, _ := json.Marshal(registerBody)

	req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Login
	loginBody := map[string]string{
		"email":    email,
		"password": password,
	}
	loginJSON, _ := json.Marshal(loginBody)

	req = httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"])
}

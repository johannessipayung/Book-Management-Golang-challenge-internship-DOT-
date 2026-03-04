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

// Cleanup user sebelum test agar tidak duplicate
func cleanupTestUser(t *testing.T, username, email string) {
	db := config.InitDB()
	userRepo := repository.NewUserRepository(db)

	if err := userRepo.DeleteByUsernameOrEmail(username, email); err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}
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
	username := fmt.Sprintf("testuser_%d", time.Now().UnixNano())
	email := fmt.Sprintf("test_%d@example.com", time.Now().UnixNano())
	password := "testpass"

	// cleanup dengan passing t
	cleanupTestUser(t, username, email)

	r := setupRouter()

	// ================= REGISTER =================
	registerBody := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	registerJSON, err := json.Marshal(registerBody)
	if err != nil {
		t.Fatalf("failed to marshal register body: %v", err)
	}

	req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// ================= LOGIN =================
	loginBody := map[string]string{
		"email":    email,
		"password": password,
	}

	loginJSON, err := json.Marshal(loginBody)
	if err != nil {
		t.Fatalf("failed to marshal login body: %v", err)
	}

	req = httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	assert.NotEmpty(t, response["token"])
}

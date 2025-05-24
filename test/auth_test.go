package test

import (
	"bytes"
	"challengeGO/handler"
	"challengeGO/model"
	"challengeGO/repository"
	"challengeGO/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

func setupAuthTest() (*gin.Engine, *gorm.DB) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	r := gin.Default()
	authHandler := handler.NewAuthHandler(service.NewUserService(repository.NewUserRepository(db)))
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	return r, db
}

func TestRegisterLogin(t *testing.T) {
	r, _ := setupAuthTest()
	user := model.User{Name: "Test", Email: "test@mail.com", Password: "pass"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	loginReq := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	loginReq.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)
	assert.Equal(t, 200, w.Code)
	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["token"])
}

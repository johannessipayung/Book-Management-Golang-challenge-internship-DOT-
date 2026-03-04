package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestJWTAuth_NoToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/protected", JWTAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 got %d", w.Code)
	}
}

func TestJWTAuth_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/protected", JWTAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 got %d", w.Code)
	}
}
func TestRoleAuthorization_Forbidden(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/admin",
		func(c *gin.Context) {
			c.Set("role", "user")
			c.Next()
		},
		RoleAuthorization("admin"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "ok"})
		},
	)

	req := httptest.NewRequest("GET", "/admin", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403 got %d", w.Code)
	}
}

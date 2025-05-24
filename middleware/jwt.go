package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("a-string-secret-at-least-256-bits-long")

type CustomClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Middleware untuk parsing dan validasi JWT
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Biasanya token dari Authorization header "Bearer token"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		// Simpan user info ke context
		c.Set("userID", claims.Subject)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// Middleware untuk cek role, boleh passing multiple role
func RoleAuthorization(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleI, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
			c.Abort()
			return
		}

		role, ok := roleI.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid role type"})
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		c.Abort()
	}
}

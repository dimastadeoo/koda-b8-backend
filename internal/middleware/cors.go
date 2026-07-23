package middleware

import (
	"net/http"
	"strings"

	"github.com/dimastadeoo/koda-b8-backend/internal/lib"
	"github.com/dimastadeoo/koda-b8-backend/internal/repo"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(sessionRepo repo.SessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}
		tokenString := parts[1]

		// Verifikasi token
		valid, userID, err := lib.VerifyToken(tokenString)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Cek session di database
		session, err := sessionRepo.FindByToken(c.Request.Context(), tokenString)
		if err != nil || session.Status != "active" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired or invalid"})
			return
		}

		c.Set("user_id", userID)
		c.Set("session_token", tokenString)
		c.Next()
	}
}
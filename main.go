package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dimastadeoo/koda-b8-backend/internal/di"
	"github.com/dimastadeoo/koda-b8-backend/internal/lib"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Koneksi database
	pool, err := lib.Conn()
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer pool.Close()

	// Seed roles
	seedRoles(context.Background(), pool)

	// Build container
	cont := di.BuildContainer(pool)

	// Router
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", cont.AuthHandler.Register)
		auth.POST("/login", cont.AuthHandler.Login)
		auth.POST("/logout", cont.AuthMiddleware, cont.AuthHandler.Logout)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(cont.AuthMiddleware)
	{
		protected.GET("/me", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"message": "protected endpoint",
			})
		})
	}

	getPort := os.Getenv("PORT_BACKEND")
	r.Run("0.0.0.0:" + getPort)
}

func seedRoles(ctx context.Context, pool *pgxpool.Pool) {
	roles := []string{"admin", "customer", "staff"}
	for _, name := range roles {
		var exists bool
		err := pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM roles WHERE name = $1)", name).Scan(&exists)
		if err != nil {
			log.Printf("Error checking role %s: %v", name, err)
			continue
		}
		if !exists {
			_, err := pool.Exec(ctx, "INSERT INTO roles (name) VALUES ($1)", name)
			if err != nil {
				log.Printf("Failed to insert role %s: %v", name, err)
			} else {
				log.Printf("Role %s seeded successfully", name)
			}
		}
	}
}
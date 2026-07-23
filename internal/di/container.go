package di

import (
	"github.com/dimastadeoo/koda-b8-backend/internal/handler"
	"github.com/dimastadeoo/koda-b8-backend/internal/middleware"
	"github.com/dimastadeoo/koda-b8-backend/internal/repo"
	"github.com/dimastadeoo/koda-b8-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	AuthHandler    *handler.AuthHandler
	AuthMiddleware gin.HandlerFunc
}

func BuildContainer(pool *pgxpool.Pool) *Container {
	userRepo := repo.NewUserRepository(pool)
	roleRepo := repo.NewRoleRepository(pool)
	profileRepo := repo.NewProfileRepository(pool)
	sessionRepo := repo.NewSessionRepository(pool)

	authService := service.NewAuthService(userRepo, roleRepo, profileRepo, sessionRepo)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleware := middleware.AuthMiddleware(sessionRepo)

	return &Container{
		AuthHandler:    authHandler,
		AuthMiddleware: authMiddleware,
	}
}
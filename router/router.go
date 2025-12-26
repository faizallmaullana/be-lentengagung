package router

import (
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/lenteng-agung/backend/database"
	handlerPkg "github.com/faizallmaullana/lenteng-agung/backend/internal/handler"
	authRepo "github.com/faizallmaullana/lenteng-agung/backend/internal/repo"
	authService "github.com/faizallmaullana/lenteng-agung/backend/internal/service"
)

// NewRouter builds a Gin engine and wires handlers using the provided DBProvider.
func NewRouter(provider database.DBProvider) (*gin.Engine, error) {
	r := gin.Default()

	// wire auth flow
	repo := authRepo.NewAuthRepo(provider)
	svc := authService.NewAuthService(repo, provider)
	handler := handlerPkg.NewAuthHandler(svc)

	api := r.Group("/api")
	{
		api.POST("/register", handler.Register())
	}

	// health check
	health := handlerPkg.NewHealthHandler(provider)
	r.GET("/health", health.Health())

	return r, nil
}

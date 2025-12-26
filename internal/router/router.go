package router

import (
	"github.com/gin-gonic/gin"

	database "github.com/faizallmaullana/lenteng-agung/backend/db/db_connection"
	handlerPkg "github.com/faizallmaullana/lenteng-agung/backend/internal/domains/handler"
	authRepo "github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"
	authService "github.com/faizallmaullana/lenteng-agung/backend/internal/domains/service"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/middleware"
)

// NewRouter builds a Gin engine and wires handlers using the provided DBProvider.
func NewRouter(provider database.DBProvider) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// wire auth flow
	repo := authRepo.NewAuthRepo(provider)

	jwtSvc := authService.NewJWTService()
	authSvc := authService.NewAuthService(repo, provider, jwtSvc)

	authHandler := handlerPkg.NewAuthHandler(authSvc)

	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/register", authHandler.Register)
	}

	// health check
	health := handlerPkg.NewHealthHandler(provider)
	r.GET("/health", health.Health())

	return r, nil
}

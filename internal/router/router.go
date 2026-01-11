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
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.Use(middleware.CORSMiddleware())

	jwtSvc := authService.NewJWTService()

	api := r.Group("/api")
	all := api.Group("")
	all.Use(middleware.JWTMiddleware(jwtSvc))

	// auth
	repo := authRepo.NewAuthRepo(provider)
	authSvc := authService.NewAuthService(repo, provider, jwtSvc)
	authHandler := handlerPkg.NewAuthHandler(authSvc)
	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	// auth approval
	approval := auth.Group("")
	approval.Use(middleware.JWTMiddlewareWithToken(jwtSvc))
	approval.POST("/approve", authHandler.Approval)

	// form
	formRepo := authRepo.NewFormRepo(provider.DB())
	formSvc := authService.NewFormService(*formRepo)
	formHandler := handlerPkg.NewFormHandler(*formSvc)

	form := all.Group("/form")
	form.POST("/create", formHandler.StartCreateForm)
	form.GET("/", formHandler.GetFormByUserID)
	form.GET("/all", formHandler.GetAllForms)

	fileUpload := all.Group("/upload")
	uploadRepo := authRepo.NewFileUploadRepo(provider.DB())
	uploadSvc := authService.NewUploadService(*uploadRepo)
	uploadHandler := handlerPkg.NewUploadHandler(*uploadSvc)

	fileUpload.POST("/:file_type", uploadHandler.UploadFile)
	fileUpload.POST("/ocr/:file_type", uploadHandler.OcrExtract)

	// health check
	health := handlerPkg.NewHealthHandler(provider)
	r.GET("/health", health.Health())

	return r, nil
}

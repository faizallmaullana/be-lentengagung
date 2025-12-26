package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faizallmaullana/lenteng-agung/backend/database"
	"github.com/faizallmaullana/lenteng-agung/backend/env"
	"github.com/faizallmaullana/lenteng-agung/backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize DB provider
	provider, err := database.NewGormProviderFromEnv()
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer provider.Close()

	// build router with wired handlers
	r, err := router.NewRouter(provider)
	if err != nil {
		log.Fatal("failed to create router: ", err)
	}

	// health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	config := env.Get()
	if err := r.Run(fmt.Sprintf(":%s", config.App.Port)); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

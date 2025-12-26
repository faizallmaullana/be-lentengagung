package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faizallmaullana/lenteng-agung/backend/env"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	config := env.Get()
	if err := router.Run(fmt.Sprintf(":%s", config.App.Port)); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

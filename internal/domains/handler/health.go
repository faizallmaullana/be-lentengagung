package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/lenteng-agung/backend/database"
)

type HealthHandler struct {
	provider database.DBProvider
}

func NewHealthHandler(provider database.DBProvider) *HealthHandler {
	return &HealthHandler{provider: provider}
}

func (h *HealthHandler) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := h.provider.DB()
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "db": false, "error": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "db": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "db": true})
	}
}

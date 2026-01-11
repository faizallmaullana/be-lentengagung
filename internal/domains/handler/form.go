package handler

import (
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/service"
	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	svc service.FormService
}

func NewFormHandler(s service.FormService) *FormHandler {
	return &FormHandler{svc: s}
}

func (h *FormHandler) StartCreateForm(c *gin.Context) {
	userID, ok := c.Get("id_user")
	if !ok {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	models, err := h.svc.CreateForm(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, models)
}

func (h *FormHandler) GetFormByUserID(c *gin.Context) {
	userID, ok := c.Get("id_user")
	if !ok {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	models, err := h.svc.GetFormByUserID(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, models)
}

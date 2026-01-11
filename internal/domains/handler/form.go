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
	

	models, err := h.svc.CreateForm()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, models)
}

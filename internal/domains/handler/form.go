package handler

import (
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/service"
)

type FormHandler struct {
	svc service.FormService
}

func NewFormHandler(s service.FormService) *FormHandler {
	return &FormHandler{svc: s}
}

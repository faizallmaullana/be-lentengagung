package service

import "github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"

type FormService struct {
	repo repo.FormRepo
}

func NewFormService(r repo.FormRepo) *FormService {
	return &FormService{repo: r}
}

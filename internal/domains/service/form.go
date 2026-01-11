package service

import (
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
)

type FormService struct {
	repo repo.FormRepo
}

func NewFormService(r repo.FormRepo) *FormService {
	return &FormService{repo: r}
}

func (s *FormService) CreateForm(userID string) (*models.RegisterPernyataan, error) {
	models, err := s.repo.CreateRequest(userID)
	if err != nil {
		return nil, err
	}
	return models, nil
}

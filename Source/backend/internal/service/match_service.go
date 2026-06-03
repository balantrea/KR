package service

import (
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/repository"
)

type MatchService struct {
	repo *repository.MatchRepository
}

func NewMatchService(repo *repository.MatchRepository) *MatchService {
	return &MatchService{repo: repo}
}

func (s *MatchService) CreateMatch(m *models.Match, userID int) error {
	return s.repo.Create(m, userID)
}

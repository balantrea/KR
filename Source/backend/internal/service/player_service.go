package service

import (
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/repository"
)

type PlayerService struct {
	repo *repository.PlayerRepository
}

func NewPlayerService(repo *repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(p *models.Player, userID int) error {
	return s.repo.Create(p, userID)
}

func (s *PlayerService) UpdatePlayerTeam(id int, teamID int, userID int) error {
	return s.repo.UpdateTeam(id, teamID, userID)
}

func (s *PlayerService) DeletePlayer(id int, userID int) error {
	return s.repo.Delete(id, userID)
}

func (s *PlayerService) GetPlayerProfiles(userID int) ([]models.PlayerProfileView, error) {
	return s.repo.GetProfiles(userID)
}

package service

import (
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/repository"
)

type TeamService struct {
	repo *repository.TeamRepository
}

func NewTeamService(repo *repository.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) CreateTeam(t *models.Team, userID int) error {
	return s.repo.Create(t, userID)
}

func (s *TeamService) GetAllTeams(userID int) ([]models.Team, error) {
	return s.repo.GetAll(userID)
}

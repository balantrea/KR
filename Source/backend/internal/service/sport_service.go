package service

import (
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/repository"
)

type SportService struct {
	repo *repository.SportRepository
}

func NewSportService(repo *repository.SportRepository) *SportService {
	return &SportService{repo: repo}
}

func (s *SportService) CreateSport(sport *models.Sport, userID int) error {
	return s.repo.Create(sport, userID)
}

func (s *SportService) GetAllSports(userID int) ([]models.Sport, error) {
	return s.repo.GetAll(userID)
}

func (s *SportService) UpdateSport(sport *models.Sport, userID int) error {
	return s.repo.Update(sport, userID)
}

func (s *SportService) DeleteSport(id int, userID int) error {
	return s.repo.Delete(id, userID)
}

func (s *SportService) GetSportTeamCount(id int, userID int) (int, error) {
	return s.repo.GetTeamCount(id, userID)
}

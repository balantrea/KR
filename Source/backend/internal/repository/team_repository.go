package repository

import (
	"database/sql"
	"sports-backend/Source/backend/internal/models"
)

type TeamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Create(t *models.Team, userID int) error {
	return r.db.QueryRow("INSERT INTO teams (name, city, sport_id, user_id) VALUES ($1, $2, $3, $4) RETURNING id", t.Name, t.City, t.SportID, userID).Scan(&t.ID)
}

func (r *TeamRepository) GetAll(userID int) ([]models.Team, error) {
	rows, err := r.db.Query("SELECT id, name, city, sport_id FROM teams WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var teams []models.Team
	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.ID, &t.Name, &t.City, &t.SportID); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	return teams, nil
}

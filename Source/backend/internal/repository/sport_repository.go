package repository

import (
	"database/sql"
	"sports-backend/Source/backend/internal/models"
)

type SportRepository struct {
	db *sql.DB
}

func NewSportRepository(db *sql.DB) *SportRepository {
	return &SportRepository{db: db}
}

func (r *SportRepository) Create(sport *models.Sport, userID int) error {
	return r.db.QueryRow("INSERT INTO sports (name, type, user_id) VALUES ($1, $2, $3) RETURNING id", sport.Name, sport.Type, userID).Scan(&sport.ID)
}

func (r *SportRepository) GetAll(userID int) ([]models.Sport, error) {
	rows, err := r.db.Query("SELECT id, name, type FROM sports WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sports []models.Sport
	for rows.Next() {
		var s models.Sport
		if err := rows.Scan(&s.ID, &s.Name, &s.Type); err != nil {
			return nil, err
		}
		sports = append(sports, s)
	}
	return sports, nil
}

func (r *SportRepository) Update(sport *models.Sport, userID int) error {
	_, err := r.db.Exec("UPDATE sports SET name = $1, type = $2 WHERE id = $3 AND user_id = $4", sport.Name, sport.Type, sport.ID, userID)
	return err
}

func (r *SportRepository) Delete(id int, userID int) error {
	_, err := r.db.Exec("DELETE FROM sports WHERE id = $1 AND user_id = $2", id, userID)
	return err
}

func (r *SportRepository) GetTeamCount(id int, userID int) (int, error) {
	var count int
	err := r.db.QueryRow("SELECT fn_get_sport_team_count($1, $2)", id, userID).Scan(&count)
	return count, err
}

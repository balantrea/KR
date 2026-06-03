package repository

import (
	"database/sql"
	"sports-backend/Source/backend/internal/models"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(p *models.Player, userID int) error {
	_, err := r.db.Exec("CALL sp_insert_player($1, $2, $3, $4, $5, $6)", p.FirstName, p.LastName, p.BirthDate, p.TeamID, p.Position, userID)
	return err
}

func (r *PlayerRepository) UpdateTeam(id int, teamID int, userID int) error {
	_, err := r.db.Exec("CALL sp_update_player_team($1, $2, $3)", id, teamID, userID)
	return err
}

func (r *PlayerRepository) Delete(id int, userID int) error {
	_, err := r.db.Exec("CALL sp_delete_player($1, $2)", id, userID)
	return err
}

func (r *PlayerRepository) GetProfiles(userID int) ([]models.PlayerProfileView, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, position, team_name, sport_name FROM vw_player_profiles WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var profiles []models.PlayerProfileView
	for rows.Next() {
		var p models.PlayerProfileView
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Position, &p.TeamName, &p.SportName); err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

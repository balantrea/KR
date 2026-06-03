package repository

import (
	"database/sql"
	"sports-backend/Source/backend/internal/models"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Create(m *models.Match, userID int) error {
	return r.db.QueryRow("INSERT INTO matches (home_team_id, away_team_id, match_date, home_score, away_score, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", m.HomeTeamID, m.AwayTeamID, m.MatchDate, m.HomeScore, m.AwayScore, userID).Scan(&m.ID)
}

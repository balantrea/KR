package models

import "time"

type Player struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	TeamID    *int      `json:"team_id"`
	Position  string    `json:"position"`
}

type PlayerProfileView struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
	TeamName  string `json:"team_name"`
	SportName string `json:"sport_name"`
}

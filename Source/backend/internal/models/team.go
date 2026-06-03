package models

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	SportID int    `json:"sport_id"`
}

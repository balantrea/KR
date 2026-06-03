package models

import "time"

type Match struct {
	ID         int       `json:"id"`
	HomeTeamID int       `json:"home_team_id"`
	AwayTeamID int       `json:"away_team_id"`
	MatchDate  time.Time `json:"match_date"`
	HomeScore  int       `json:"home_score"`
	AwayScore  int       `json:"away_score"`
}

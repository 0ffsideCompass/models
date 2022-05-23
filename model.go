package models

import "time"

// ClubSeasonDataJSONRequest represents the request for the endpoint /club/season
type ClubSeasonDataJSONRequest struct {
	League string         `json:"league"`
	Data   []ModelFixture `json:"data"`
}

// ModelFixture represents a game between two teams
type ModelFixture struct {
	TeamOne string `json:"teamOne"`
	TeamTwo string `json:"teamTwo"`
}

// ClubSeasonDataJSONResponse response for endpoint /club/season
type ClubSeasonDataJSONResponse struct {
	HistoricalData []HistoricalData `json:"data"`
}

// HistoricalData respresnt the hitorical games between two games
type HistoricalData struct {
	TeamsID  string        `json:"id"`
	TeamOne  string        `json:"teamOne"`
	TeamTwo  string        `json:"teamTwo"`
	Fixtures []ClubFixture `json:"fixtures"`
}

// ClubFixture represents the details of a game between two teams
type ClubFixture struct {
	Div              string    `json:"div"`
	Date             time.Time `json:"date"`
	HomeTeam         string    `json:"homeTeam"`
	AwayTeam         string    `json:"awayTeam"`
	FullTimeHomeGoal string    `json:"fullTimeHomeGoal"`
	FullTimeAwayGoal string    `json:"fullTimeAwayGoal"`
	Winner           string    `json:"winner"`
	HalfTimeHomeGoal string    `json:"halfTimeHomeGoal"`
	HalfTimeAwayGoal string    `json:"halfTimeAwayGoal"`
	HalfTimeWinner   string    `json:"halfTimeWinner"`
	Attendance       string    `json:"attendance"`
	Referee          string    `json:"referee"`
	Season           int32     `json:"season"`
}

// WorldCupDataJSON represents all world cup data teams
type WorldCupDataJSON struct {
	TotalGoals          int32  `json:"total_goals"`
	HalfTimeGoals       int32  `json:"half_time_goals"`
	TeamGames           []Team `json:"team_games"`
	MostSuccessfulTeams []Team `json:"most_successful_teams"`
}

// Team represents a world cup participated team
type Team struct {
	Nation   string `json:"nation"`
	Amount   int    `json:"amount"`
	Position int    `json:"position"`
}

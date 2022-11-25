package models

import "time"

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	League    string    `json:"league"`
	HTML      string    `json:"html"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []string  `json:"tags"`
	Data      []Data    `json:"data"`
}

type Articles struct {
	Data []Article `json:"data"`
}

type Data struct {
	TeamOne               string           `json:"teamOne"`
	TeamTwo               string           `json:"teamTwo"`
	TeamOneWins           int              `json:"teamOneWins"`
	TeamTwoWins           int              `json:"teamTwoWins"`
	Draws                 int              `json:"draws"`
	TotalGames            int              `json:"totalGames"`
	TeamOneWinsPercentage int              `json:"teamOneWinsPercentage"`
	TeamTwoWinsPercentage int              `json:"teamTwoWinsPercentage"`
	DrawsPercentage       int              `json:"drawsPercentage"`
	HistoricalGames       []HistoricalGame `json:"historicalData"`
}

type HistoricalGame struct {
	Div              string    `json:"Div"`
	Date             time.Time `json:"Date"`
	HomeTeam         string    `json:"HomeTeam"`
	AwayTeam         string    `json:"AwayTeam"`
	FullTimeHomeGoal string    `json:"FullTimeHomeGoal"`
	FullTimeAwayGoal string    `json:"FullTimeAwayGoal"`
	Winner           string    `json:"Winner"`
	HalfTimeHomeGoal string    `json:"HalfTimeHomeGoal"`
	HalfTimeAwayGoal string    `json:"HalfTimeAwayGoal"`
	HalfTimeWinner   string    `json:"HalfTimeWinner"`
	Attendance       string    `json:"Attendance"`
	Referee          string    `json:"Referee"`
	Season           int32     `json:"season"`
}

type ArticleImportanceList struct {
	PositionOne   ArticleImportance `json:"positionOne"`
	PositionTwo   ArticleImportance `json:"positionTwo"`
	PositionThree ArticleImportance `json:"positionThree"`
	PositionFour  ArticleImportance `json:"positionFour"`
	PositionFive  ArticleImportance `json:"positionFive"`
	PositionSix   ArticleImportance `json:"positionSix"`
	PositionSeven ArticleImportance `json:"positionSeven"`
}

type ArticleImportance struct {
	ArticleID    string `bson:"article_id"`
	ArticleTitle string `bson:"article_title"`
}

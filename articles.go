package models

import "time"

type Article struct {
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
	TeamOne         string           `json:"teamOne"`
	TeamTwo         string           `json:"teamTwo"`
	HistoricalGames []HistoricalGame `json:"historicalData"`
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

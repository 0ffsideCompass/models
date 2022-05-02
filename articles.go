package models

import "time"

type Fixture struct {
	Title     string    `json:"title"`
	League    string    `json:"league"`
	HTML      string    `json:"html"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []string  `json:"tags"`
}

type FixturesByLeague struct {
	Data []Fixture `json:"data"`
}

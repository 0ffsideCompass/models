package models

import "time"

type Article struct {
	Title     string    `json:"title"`
	League    string    `json:"league"`
	HTML      string    `json:"html"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []string  `json:"tags"`
}

type Articles struct {
	Data []Article `json:"data"`
}

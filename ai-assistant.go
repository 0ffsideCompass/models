package models

type AskRequest struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

type AskResponse struct {
	Text string `json:"text"`
}

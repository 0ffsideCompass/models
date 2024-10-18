package models

type AskRequest struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

type AskResponse struct {
	Text    string   `json:"text"`
	Metadat []string `json:"metadata"`
}

type PromptHistory struct {
	Prompts []PromptReqResponse `json:"prompts"`
}

type PromptReqResponse struct {
	Prompt   string `json:"prompt"`
	Response string `json:"response"`
}

type AskCoverageRequest struct {
	Text      string `json:"text"`
	FixtureID string `json:"fixture_id"`
	ChatID    string `json:"chat_id"`
}

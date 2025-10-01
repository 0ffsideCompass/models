package models

// DataWarehouseCreateArticleRequest represents the request to create an article
type DataWarehouseCreateArticleRequest struct {
	ExternalID string   `json:"external_id"`
	Title      string   `json:"title"`
	URL        string   `json:"url"`
	Tags       []string `json:"tags"`
}

// DataWarehouseArticlesResponse represents the API response for multiple articles
type DataWarehouseArticlesResponse struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
	Page     int       `json:"page,omitempty"`
	Limit    int       `json:"limit,omitempty"`
	Error    string    `json:"error,omitempty"`
}

// DataWarehouseCreatePodcastRequest represents the request to create a podcast
type DataWarehouseCreatePodcastRequest struct {
	Title      string   `json:"title"`
	ExternalID string   `json:"external_id"`
	URL        string   `json:"url"`
	Tags       []string `json:"tags"`
}

// DataWarehousePodcastsResponse represents the API response for multiple podcasts
type DataWarehousePodcastsResponse struct {
	Podcasts []Podcast `json:"podcasts"`
	Total    int       `json:"total"`
	Page     int       `json:"page,omitempty"`
	Limit    int       `json:"limit,omitempty"`
	Error    string    `json:"error,omitempty"`
}

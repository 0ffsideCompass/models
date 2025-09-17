package models

import "time"

// DataWarehouseArticle represents the domain model for an article
type DataWarehouseArticle struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DataWarehouseCreateArticleRequest represents the request to create an article
type DataWarehouseCreateArticleRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
}

// DataWarehouseUpdateArticleRequest represents the request to update an article
type DataWarehouseUpdateArticleRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// DataWarehouseArticleResponse represents the API response for an article
type DataWarehouseArticleResponse struct {
	Article *DataWarehouseArticle `json:"article,omitempty"`
	Error   string                `json:"error,omitempty"`
}

// DataWarehouseArticlesResponse represents the API response for multiple articles
type DataWarehouseArticlesResponse struct {
	Articles []DataWarehouseArticle `json:"articles"`
	Total    int                    `json:"total"`
	Page     int                    `json:"page,omitempty"`
	Limit    int                    `json:"limit,omitempty"`
	Error    string                 `json:"error,omitempty"`
}

// DataWarehousePodcast represents the domain model for a podcast
type DataWarehousePodcast struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DataWarehouseCreatePodcastRequest represents the request to create a podcast
type DataWarehouseCreatePodcastRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
}

// DataWarehouseUpdatePodcastRequest represents the request to update a podcast
type DataWarehouseUpdatePodcastRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// DataWarehousePodcastResponse represents the API response for a podcast
type DataWarehousePodcastResponse struct {
	Podcast *DataWarehousePodcast `json:"podcast,omitempty"`
	Error   string                `json:"error,omitempty"`
}

// DataWarehousePodcastsResponse represents the API response for multiple podcasts
type DataWarehousePodcastsResponse struct {
	Podcasts []DataWarehousePodcast `json:"podcasts"`
	Total    int                    `json:"total"`
	Page     int                    `json:"page,omitempty"`
	Limit    int                    `json:"limit,omitempty"`
	Error    string                 `json:"error,omitempty"`
}
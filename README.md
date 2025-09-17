# Offside Compass Models Package

This package contains all the data models and structures used across the Offside Compass ecosystem, including the main application, data warehouse integration, and various APIs.

## Package Structure

The models are organized into several files based on their domain and usage:

### Core Models

#### `model.go`
Contains the primary fixture and historical data models:
- `ClubSeasonDataJSONRequest/Response` - Season data for clubs
- `ModelFixture` - Basic fixture representation
- `HistoricalData` - Historical match records between teams
- `ClubFixture` - Detailed fixture information

#### `articles.go`
Article management and content models:
- `Article` - Main article structure with HTML content, tags, and metadata
- `Articles` - Collection wrapper for articles
- `Data` - Match-specific data including team statistics
- `TeamStats` - Comprehensive team performance statistics

#### `podcast.go`
Podcast content models:
- `Podcast` - Podcast episodes with audio URLs and scripts
- `PodcastArchiveRequests` - Archive request handling

### API Integration Models

#### `api-football.go`
External football API integration models:
- `TodayFixture` - Daily fixture collections
- `LeagueFixture` - League-specific fixtures
- `Fixture` - Match details including venue, teams, and statistics
- `LastGame` - Previous match results
- `Stats` - Team performance metrics

#### `ai-assistant.go`
AI chat and assistant functionality:
- `AskRequest/Response` - Chat interaction models
- `ChatHistory` - Conversation persistence
- `PromptHistory` - Prompt tracking and management

### Data Warehouse Integration

#### `warehouse.go`
Models for data warehouse synchronization with `DataWarehouse` prefix:
- `DataWarehouseArticle` - Article synchronization model
- `DataWarehousePodcast` - Podcast synchronization model
- Associated request/response structures for CRUD operations

## Usage

Import the package in your Go code:

```go
import "github.com/yourusername/offsidecompass/models"
```

Example usage:

```go
// Create a new article
article := models.Article{
    ID:        "unique-id",
    Title:     "Match Preview",
    League:    "Premier League",
    HTML:      "<p>Content here</p>",
    CreatedAt: time.Now(),
    Tags:      []string{"preview", "analysis"},
}

// Work with fixtures
fixture := models.Fixture{
    HomeTeam: "Team A",
    AwayTeam: "Team B",
    Stadium:  "Stadium Name",
    City:     "London",
}
```

## Model Categories

1. **Content Models** - Articles, Podcasts
2. **Match Data** - Fixtures, Teams, Statistics
3. **API Models** - Request/Response structures
4. **AI/Chat Models** - Conversation and prompt management
5. **Data Warehouse Models** - Synchronization structures

## Dependencies

- `time` package for timestamps
- `go.mongodb.org/mongo-driver` for MongoDB integration (in ai-assistant.go)

## Notes

- All models use JSON tags for serialization
- Validation tags are included where appropriate
- Time fields use `time.Time` for proper date handling
- IDs are typically strings for flexibility across different storage systems
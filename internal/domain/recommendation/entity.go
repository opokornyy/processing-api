package recommendation

import (
	"time"
)

type RecommendationID string

// Recommendation represents an AI-generated recommendation
type Recommendation struct {
	ID          RecommendationID
	WeatherID   string
	Title       string
	Description string
	CreatedAt   time.Time
}

func NewRecommendation(
	id RecommendationID,
	weatherID string,
	title string,
	description string,
) *Recommendation {
	return &Recommendation{
		ID:          id,
		WeatherID:   weatherID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

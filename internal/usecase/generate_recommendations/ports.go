package generaterecommendations

import (
	"processing-api/internal/domain/recommendation"
	"processing-api/internal/domain/weather"
)

// WeatherRepository defines the interface for reading weather data
type WeatherRepository interface {
	GetByID(id weather.WeatherID) (*weather.WeatherData, error)
}

// AIService defines the interface for AI recommendation generation based on weather data
type AIService interface {
	GenerateRecommendation(weather *weather.WeatherData) (*recommendation.Recommendation, error)
}

// RecommendationRepository defines the interface for recommendation persistence
type RecommendationRepository interface {
	Save(rec *recommendation.Recommendation) error
}

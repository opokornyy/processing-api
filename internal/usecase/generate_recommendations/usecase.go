package generaterecommendations

import (
	"context"

	"processing-api/internal/domain/weather"
)

type GenerateRecommendationsUseCase struct {
	weatherRepo        WeatherRepository
	aiService          AIService
	recommendationRepo RecommendationRepository
}

func NewGenerateRecommendationsUseCase(
	weatherRepo WeatherRepository,
	aiService AIService,
	recommendationRepo RecommendationRepository,
) *GenerateRecommendationsUseCase {
	return &GenerateRecommendationsUseCase{
		weatherRepo:        weatherRepo,
		aiService:          aiService,
		recommendationRepo: recommendationRepo,
	}
}

// Execute generates AI recommendations for given weather data
func (uc *GenerateRecommendationsUseCase) Execute(ctx context.Context, weatherID weather.WeatherID) error {
	// TODO: 1. Get weather data by ID from repository
	// TODO: 2. Call AI service to generate recommendation
	// TODO: 3. Save recommendation to repository

	return nil
}

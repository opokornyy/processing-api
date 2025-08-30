package processweatherworkflow

import (
	"processing-api/internal/domain/weather"
	fetchweather "processing-api/internal/usecase/fetch_weather"
	generaterecommendations "processing-api/internal/usecase/generate_recommendations"
)

// ProcessWeatherWorkflowUseCase orchestrates the complete weather processing workflow
type ProcessWeatherWorkflowUseCase struct {
	fetchWeatherUseCase            *fetchweather.FetchWeatherUseCase
	generateRecommendationsUseCase *generaterecommendations.GenerateRecommendationsUseCase
	notificationService            NotificationService
}

// NewProcessWeatherWorkflowUseCase creates a new instance of the workflow usecase
func NewProcessWeatherWorkflowUseCase(
	fetchWeatherUseCase *fetchweather.FetchWeatherUseCase,
	generateRecommendationsUseCase *generaterecommendations.GenerateRecommendationsUseCase,
	notificationService NotificationService,
) *ProcessWeatherWorkflowUseCase {
	return &ProcessWeatherWorkflowUseCase{
		fetchWeatherUseCase:            fetchWeatherUseCase,
		generateRecommendationsUseCase: generateRecommendationsUseCase,
		notificationService:            notificationService,
	}
}

// Execute runs the complete workflow: fetch weather → generate recommendations → send notification
func (uc *ProcessWeatherWorkflowUseCase) Execute(location weather.Location) error {
	// TODO: 1. Fetch weather data for location
	// TODO: 2. Generate AI recommendations based on weather
	// TODO: 3. Send recommendation via notification service

	return nil
}

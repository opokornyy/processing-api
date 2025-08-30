package weatherapi

import (
	"context"
	"fmt"

	"processing-api/internal/domain/weather"

	"github.com/rs/zerolog/log"
)

type WeatherAPI struct{}

func NewWeatherAPI() *WeatherAPI {
	return &WeatherAPI{}
}

func (wa *WeatherAPI) GetCurrentWeather(ctx context.Context, location weather.Location) (*weather.WeatherData, error) {
	if location.City == "" || location.Country == "" {
		log.Error().Msgf("openWeather: location.city and location.country are required: %v", location)
		return nil, fmt.Errorf("openWeather: location.city and location.country are required")
	}

	// Mock weather data for OpenWeather API
	mockWeather := weather.NewWeatherData(
		weather.WeatherID("ow-"+location.City+"-"+location.Country),
		location,
		weather.Temperature{Celsius: 22.5},
		65.0,
	)
	return mockWeather, nil
}

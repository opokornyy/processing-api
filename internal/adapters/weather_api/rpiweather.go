package weatherapi

import (
	"context"
	"fmt"

	"processing-api/internal/domain/weather"

	"github.com/rs/zerolog/log"
)

type WeatherRPI struct{}

func NewWeatherRPI() *WeatherRPI {
	return &WeatherRPI{}
}

func (wr *WeatherRPI) GetCurrentWeather(ctx context.Context, location weather.Location) (*weather.WeatherData, error) {
	if location.City == "" || location.Country == "" {
		log.Error().Msgf("rpiWeather: location.city and location.country are required: %v", location)
		return nil, fmt.Errorf("rpiWeather: location.city and location.country are required")
	}

	// Mock weather data for RPI Weather API
	mockWeather := weather.NewWeatherData(
		weather.WeatherID("rpi-"+location.City+"-"+location.Country),
		location,
		weather.Temperature{Celsius: 18.3},
		72.5,
	)
	return mockWeather, nil
}

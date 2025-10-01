package weatherapi

import (
	"context"

	"processing-api/internal/domain/weather"
)

const openWeatherName = "OpenWeather"

type WeatherAPI struct{}

func NewWeatherAPI() *WeatherAPI {
	return &WeatherAPI{}
}

func (wa *WeatherAPI) GetCurrentWeather(ctx context.Context, location weather.Location) (*weather.WeatherData, error) {
	if location.City == "" || location.Country == "" {
		return nil, WeatherApiError{
			ApiName: openWeatherName,
			Message: "city and country are required",
			Err:     ErrMissingArgs,
		}
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

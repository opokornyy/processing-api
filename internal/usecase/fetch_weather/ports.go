package fetchweather

import (
	"context"

	"processing-api/internal/domain/weather"
)

// WeatherAPI defines the interface for external weather services
type WeatherAPI interface {
	GetCurrentWeather(ctx context.Context, location weather.Location) (*weather.WeatherData, error)
}

// WeatherRepository defines the interface for weather data persistence
type WeatherRepository interface {
	Save(ctx context.Context, weather *weather.WeatherData) error
	GetLatestByLocation(ctx context.Context, location weather.Location) (*weather.WeatherData, error)
}

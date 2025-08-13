package fetchweather

import (
	"processing-api/internal/domain/weather"
)

// WeatherAPI defines the interface for external weather services
type WeatherAPI interface {
	GetCurrentWeather(location weather.Location) (*weather.WeatherData, error)
}

// WeatherRepository defines the interface for weather data persistence
type WeatherRepository interface {
	Save(weather *weather.WeatherData) error
	GetLatestByLocation(location weather.Location) (*weather.WeatherData, error)
}

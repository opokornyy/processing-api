package fetchweather

import (
	"processing-api/internal/domain/weather"
)

type FetchWeatherUseCase struct {
	weatherApi  WeatherAPI
	weatherRepo WeatherRepository
}

func NewFetchWeatherUseCase(weatherRepo WeatherRepository) *FetchWeatherUseCase {
	return &FetchWeatherUseCase{
		weatherRepo: weatherRepo,
	}
}

// Execute fetches weather data for a location and stores it
func (uc *FetchWeatherUseCase) Execute(location weather.Location) (*weather.WeatherData, error) {
	// TODO: 1. Fetch weather data from external API
	// TODO: 2. Create weather entity with proper ID
	// TODO: 3. Save to repository
	// TODO: 4. Return the weather data

	return nil, nil
}

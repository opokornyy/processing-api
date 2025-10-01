package fetchweather

import (
	"context"

	"processing-api/internal/domain/weather"

	"github.com/rs/zerolog/log"
)

const featchWeatherName = "featchWeatherUseCase"

type FetchWeatherUseCase struct {
	weatherApi  WeatherAPI
	weatherRepo WeatherRepository
}

func NewFetchWeatherUseCase(api WeatherAPI, weatherRepo WeatherRepository) *FetchWeatherUseCase {
	return &FetchWeatherUseCase{
		weatherApi:  api,
		weatherRepo: weatherRepo,
	}
}

// Execute fetches weather data for a location and stores it
func (uc *FetchWeatherUseCase) Execute(ctx context.Context, location weather.Location) (*weather.WeatherData, error) {
	// 1. Fetch weather data from external API
	weatherData, err := uc.weatherApi.GetCurrentWeather(ctx, location)
	if err != nil {
		log.Err(err).Msgf("%s: fetching location: %s", featchWeatherName, location)
		return nil, err
	}

	// 2. Save to repository
	if err := uc.weatherRepo.Save(ctx, weatherData); err != nil {
		log.Err(err).Msgf("%s: saving data: %s", featchWeatherName, *weatherData)
		return nil, err
	}

	// 3. Return the weather data
	return weatherData, nil
}

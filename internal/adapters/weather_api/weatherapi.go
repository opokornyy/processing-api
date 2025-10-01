// Package weatherapi provides adapters for integrating with external weather API services.
// It implements the WeatherAPI interface from the domain layer to fetch current weather data
// from various third-party weather providers such as weatherapi.com
package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"processing-api/config"
	"processing-api/internal/domain/weather"

	"github.com/rs/zerolog/log"
)

const (
	openWeatherName        = "OpenWeather"
	currentWeatherEndpoint = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"
)

type WeatherAPIResponse struct {
	Location WeatherAPILocation
	Current  WeatherAPICurrent
}

type WeatherAPILocation struct {
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type WeatherAPICurrent struct {
	TempC     float64 `json:"temp_c"`
	Humidity  float64 `json:"humidity"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
}

type WeatherAPI struct {
	config config.Config
}

func NewWeatherAPI(config config.Config) *WeatherAPI {
	return &WeatherAPI{
		config: config,
	}
}

func (wa *WeatherAPI) GetCurrentWeather(ctx context.Context, location weather.Location) (*weather.WeatherData, error) {
	log.Info().
		Str("city", location.City).
		Str("country", location.Country).
		Msg("fetching weather data")

	if location.City == "" || location.Country == "" {
		log.Error().Msg("missing required location parameters")
		return nil, WeatherApiError{
			ApiName: openWeatherName,
			Message: "city and country are required",
			Err:     ErrMissingArgs,
		}
	}

	// Build the request URL with the city name
	url := fmt.Sprintf(currentWeatherEndpoint, wa.config.WeatherAPIKey, location.City)
	log.Debug().Str("url", url).Msg("built API request URL")

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create HTTP request")
		return nil, WeatherApiError{
			ApiName: openWeatherName,
			Message: "failed to create request",
			Err:     err,
		}
	}

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("failed to execute HTTP request")
		return nil, WeatherApiError{
			ApiName: openWeatherName,
			Message: "failed to fetch weather data",
			Err:     err,
		}
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		log.Error().
			Int("status_code", resp.StatusCode).
			Str("city", location.City).
			Msg("unexpected status code from weather API")
		return nil, WeatherApiError{
			ApiName: openWeatherName,
			Message: fmt.Sprintf("unexpected status code: %d", resp.StatusCode),
			Err:     fmt.Errorf("status code %d", resp.StatusCode),
		}
	}

	log.Debug().Msg("received successful response from weather API")

	// Decode the response
	var apiResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Error().Err(err).Msg("failed to decode JSON response")
		return nil, WeatherApiError{
			ApiName: openWeatherName,
			Message: "failed to decode response",
			Err:     err,
		}
	}

	// Convert API response to domain model
	weatherData := weather.NewWeatherData(
		weather.WeatherID(fmt.Sprintf("wa-%s-%s", apiResponse.Location.Name, apiResponse.Location.Country)),
		weather.Location{
			City:    apiResponse.Location.Name,
			Country: apiResponse.Location.Country,
		},
		weather.Temperature{Celsius: apiResponse.Current.TempC},
		apiResponse.Current.Humidity,
	)

	log.Info().
		Str("city", apiResponse.Location.Name).
		Str("country", apiResponse.Location.Country).
		Float64("temp_c", apiResponse.Current.TempC).
		Float64("humidity", apiResponse.Current.Humidity).
		Msg("successfully fetched and converted weather data")

	return weatherData, nil
}

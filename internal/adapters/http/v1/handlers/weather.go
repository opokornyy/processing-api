package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"processing-api/internal/adapters/repository"
	weatherapi "processing-api/internal/adapters/weather_api"
	"processing-api/internal/domain/weather"
	fetchweather "processing-api/internal/usecase/fetch_weather"

	"github.com/rs/zerolog/log"
)

type WeatherHandler struct {
	fetchWeatherUC *fetchweather.FetchWeatherUseCase
}

func NewWeatherHandler(fetchWeatherUC *fetchweather.FetchWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{
		fetchWeatherUC: fetchWeatherUC,
	}
}

type WeatherResponse struct {
	ID          string  `json:"id"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

func (wh *WeatherHandler) GetWeatherForLocation(w http.ResponseWriter, r *http.Request) {
	weatherLocation, err := parseLocationParams(r)
	if err != nil {
		log.Error().Msg("Failed to parse the location")
		writeErrorResponse(w, http.StatusBadRequest, "location_invalid", err.Error())
		return
	}
	log.Info().Msgf(
		"Get /api/v1/weather/%s/%s request",
		weatherLocation.Country,
		weatherLocation.City,
	)

	weatherData, err := wh.fetchWeatherUC.Execute(r.Context(), *weatherLocation)
	if err != nil {
		handleFetchWeatherErr(err, w, weatherLocation)
		return
	}

	response := WeatherResponse{
		ID:          string(weatherData.ID),
		City:        weatherData.Location.City,
		Country:     weatherData.Location.Country,
		Temperature: weatherData.Temperature.Celsius,
		Humidity:    weatherData.Humidity,
		Timestamp:   weatherData.Timestamp.UTC().Format(time.RFC1123),
	}

	data, err := json.Marshal(response)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "serialization_error", fmt.Sprintf("Failed to serialize response: %v", err))
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleFetchWeatherErr(err error, w http.ResponseWriter, weatherLocation *weather.Location) {
	var statusCode int
	var errorCode string
	var userMessage string

	switch {
	case errors.Is(err, repository.ErrNotFound):
		statusCode = http.StatusNotFound
		errorCode = "location_not_found"
		userMessage = fmt.Sprintf("No weather data could be found for %s, %s.", weatherLocation.City, weatherLocation.Country)

	case errors.Is(err, repository.ErrEmptyData):
		statusCode = http.StatusBadRequest
		errorCode = "missing_data"
		userMessage = "Can not store empty data"

	case errors.Is(err, weatherapi.ErrMissingArgs):
		statusCode = http.StatusBadRequest
		errorCode = "invalid_request"
		userMessage = "city and country are required."

	case errors.Is(err, weatherapi.ErrFetchingWeather):
		statusCode = http.StatusBadRequest
		errorCode = "external_api_error"
		userMessage = "There was a problem fetching data from the external weather service."

	default:
		statusCode = http.StatusInternalServerError
		errorCode = "internal_error"
		userMessage = "An unexpected error occurred. Please try again later."
		log.Err(err).Msg("Unknown error in weather handler")
	}

	writeErrorResponse(w, statusCode, errorCode, userMessage)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, errorType, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResp := ErrorResponse{
		Error:   errorType,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(errorResp); err != nil {
		log.Error().Err(err).Msg("Failed to encode error response")
	}
}

func parseLocationParams(r *http.Request) (*weather.Location, error) {
	city := r.PathValue("city")
	country := r.PathValue("country")
	if city == "" || country == "" {
		return nil, fmt.Errorf("failed to parse the location param")
	}

	return &weather.Location{
		City:    city,
		Country: country,
	}, nil
}

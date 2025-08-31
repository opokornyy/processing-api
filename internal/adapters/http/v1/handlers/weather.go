package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
		wh.writeErrorResponse(w, http.StatusBadRequest, "location_invalid", err.Error())
		return
	}
	log.Info().Msgf(
		"Get /api/v1/weather/%s/%s request",
		weatherLocation.Country,
		weatherLocation.City,
	)

	weatherData, err := wh.fetchWeatherUC.Execute(r.Context(), *weatherLocation)
	if err != nil {
		log.Err(err).Msg("Failed to fetch the weather data")
		wh.writeErrorResponse(w, http.StatusBadGateway, "weather_unvailable", fmt.Sprintf("Weather not found for location: %s/%s", weatherLocation.Country, weatherLocation.City))
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
		wh.writeErrorResponse(w, 500, err.Error(), "Failed to serialize")
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (wh *WeatherHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, errorType, message string) {
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

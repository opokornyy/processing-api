package weather

import (
	"time"
)

type WeatherID string

// Location represents a geographic location
type Location struct {
	City    string
	Country string
}

// Temperature represents temperature in Celsius
type Temperature struct {
	Celsius float64
}

// WeatherData represents the main weather entity
type WeatherData struct {
	ID          WeatherID
	Location    Location
	Temperature Temperature
	Humidity    float64
	Timestamp   time.Time
}

func NewWeatherData(
	id WeatherID,
	location Location,
	temperature Temperature,
	humidity float64,
) *WeatherData {
	return &WeatherData{
		ID:          id,
		Location:    location,
		Temperature: temperature,
		Humidity:    humidity,
		Timestamp:   time.Now(),
	}
}

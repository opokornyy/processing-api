package weather

import (
	"fmt"
	"strconv"
	"time"
)

type WeatherID string

// Location represents a geographic location
type Location struct {
	City    string
	Country string
}

func (l Location) String() string {
	return fmt.Sprintf("city: %s, country: %s", l.City, l.Country)
}

// Temperature represents temperature in Celsius
type Temperature struct {
	Celsius float64
}

func (t Temperature) String() string {
	return strconv.FormatFloat(t.Celsius, 'f', 2, 64)
}

// WeatherData represents the main weather entity
type WeatherData struct {
	ID          WeatherID
	Location    Location
	Temperature Temperature
	Humidity    float64
	Timestamp   time.Time
}

func (wd WeatherData) String() string {
	return fmt.Sprintf("ID: %s, Location: %s, Temperature: %s, Humidity: %.1f%%, Timestamp: %s",
		wd.ID, wd.Location, wd.Temperature, wd.Humidity,
		wd.Timestamp.Format(time.RFC3339))
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
		Timestamp:   time.Now().UTC(),
	}
}

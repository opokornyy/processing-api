package mock

import (
	"context"
	"errors"
	"sync"

	"processing-api/internal/domain/weather"
)

// MockWeatherRepository implements WeatherRepository interface for testing
type MockWeatherRepository struct {
	mu         sync.RWMutex
	data       map[weather.WeatherID]*weather.WeatherData
	byLocation map[string]*weather.WeatherData // key: "city,country"
}

func NewMockWeatherRepository() *MockWeatherRepository {
	repo := &MockWeatherRepository{
		data:       make(map[weather.WeatherID]*weather.WeatherData),
		byLocation: make(map[string]*weather.WeatherData),
	}

	// Add some sample data
	repo.seedData()
	return repo
}

// seedData adds sample weather data for testing/development
func (m *MockWeatherRepository) seedData() {
	// Prague weather data
	pragueWeather := weather.NewWeatherData(
		weather.WeatherID("weather-prague-001"),
		weather.Location{City: "Prague", Country: "CZ"},
		weather.Temperature{Celsius: 22.5},
		65.0, // humidity percentage
	)
	m.Save(context.Background(), pragueWeather)

	// London weather data
	londonWeather := weather.NewWeatherData(
		weather.WeatherID("weather-london-001"),
		weather.Location{City: "London", Country: "UK"},
		weather.Temperature{Celsius: 18.0},
		78.0, // humidity percentage
	)
	m.Save(context.Background(), londonWeather)

	// New York weather data
	nyWeather := weather.NewWeatherData(
		weather.WeatherID("weather-ny-001"),
		weather.Location{City: "New York", Country: "US"},
		weather.Temperature{Celsius: 25.0},
		55.0, // humidity percentage
	)
	m.Save(context.Background(), nyWeather)

	// Berlin weather data
	berlinWeather := weather.NewWeatherData(
		weather.WeatherID("weather-berlin-001"),
		weather.Location{City: "Berlin", Country: "DE"},
		weather.Temperature{Celsius: 20.0},
		70.0, // humidity percentage
	)
	m.Save(context.Background(), berlinWeather)
}

// Save stores weather data in memory
func (m *MockWeatherRepository) Save(ctx context.Context, weatherData *weather.WeatherData) error {
	if weatherData == nil {
		return errors.New("weather data cannot be nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[weatherData.ID] = weatherData
	locationKey := m.getLocationKey(weatherData.Location)
	m.byLocation[locationKey] = weatherData

	return nil
}

// GetLatestByLocation retrieves the latest weather data for a location
func (m *MockWeatherRepository) GetLatestByLocation(ctx context.Context, location weather.Location) (*weather.WeatherData, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	locationKey := m.getLocationKey(location)
	weatherData, exists := m.byLocation[locationKey]
	if !exists {
		return nil, errors.New("weather data not found for location")
	}

	return weatherData, nil
}

// Helper method to create location key
func (m *MockWeatherRepository) getLocationKey(location weather.Location) string {
	return location.City + "," + location.Country
}

// Additional helper methods for testing

// Clear removes all data from the mock repository
func (m *MockWeatherRepository) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = make(map[weather.WeatherID]*weather.WeatherData)
	m.byLocation = make(map[string]*weather.WeatherData)
}

// Count returns the number of weather records stored
func (m *MockWeatherRepository) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}

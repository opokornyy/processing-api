package weatherapi

import (
	"errors"
	"fmt"
)

var (
	ErrFetchingWeather = errors.New("failed fetching weather")
	ErrMissingArgs     = errors.New("missing required arguments")
)

type WeatherApiError struct {
	ApiName string
	Message string
	Err     error
}

func (e WeatherApiError) Error() string {
	return fmt.Sprintf("%s: %s: %v", e.ApiName, e.Message, e.Err)
}

func (e WeatherApiError) Unwrap() error {
	return e.Err
}

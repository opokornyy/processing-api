package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"processing-api/config"
	"processing-api/internal/adapters/http/v1/handlers"
	"processing-api/internal/adapters/repository/mock"
	weatherapi "processing-api/internal/adapters/weather_api"
	fetchweather "processing-api/internal/usecase/fetch_weather"

	"github.com/rs/zerolog/log"
)

func Run(cfg config.Config) {
	log.Info().Msg("Starting the application")

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", cfg.Port),
		Handler: mux,
	}

	weatherAPi := weatherapi.NewWeatherAPI()
	weatherRepo := mock.NewMockWeatherRepository()
	fetchWeatherUC := fetchweather.NewFetchWeatherUseCase(weatherAPi, weatherRepo)

	// Create weather handlers
	weatherHandler := handlers.NewWeatherHandler(fetchWeatherUC)

	// Register weather routes
	mux.HandleFunc("GET /api/v1/weather/{country}/{city}", weatherHandler.GetWeatherForLocation)

	go func() {
		log.Info().Str("addr", server.Addr).Msg("HTTP server starting")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited")
}

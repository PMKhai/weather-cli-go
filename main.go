package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define command-line flags
	city := flag.String("city", "Ho Chi Minh City", "City naming")

	// Parse the command-line flags
	flag.Parse()

	// Load environment file before creating new WeatherService
	if err := loadEnvFile(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	service := NewWeatherService()

	// Get coordinates for the provided city
	coordinates, err := service.getCoordinates(*city)
	if err != nil {
		log.Fatalf("Failed to get coordinates: %v", err)
	}

	// Retrieve weather information for all coordinates
	for _, coordinate := range coordinates {
		weather, err := service.getWeather(coordinate.Lat, coordinate.Lon)
		if err != nil {
			log.Printf("Error retrieving weather for [%s]: %v\n", coordinate.Name, err)
			continue
		}
		fmt.Printf("Weather in %s: %+v\n", coordinate.Name, weather.Main)
	}
}

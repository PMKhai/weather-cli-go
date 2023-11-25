package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CoordinatesResponse struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type Weather struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Id          int    `json:"id"`
	Main        string `json:"main"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

type WeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type WeatherService struct {
	API_KEY                string
	API_GEO_V1             string
	API_CURRENT_WEATHER_V2 string
	client                 *http.Client
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		API_KEY:                getEnvOrDefault("API_KEY", ""),
		API_GEO_V1:             getEnvOrDefault("API_GEO_V1", ""),
		API_CURRENT_WEATHER_V2: getEnvOrDefault("API_CURRENT_WEATHER_V2", ""),
		client:                 &http.Client{},
	}
}

func (ws *WeatherService) getWeather(lat float64, lon float64) (*WeatherResponse, error) {
	params := "?lat=" + fmt.Sprintf("%.4f", lat) + "&lon=" + fmt.Sprintf("%.4f", lon) + "&appid=" + ws.API_KEY
	resp, err := ws.client.Get(ws.API_CURRENT_WEATHER_V2 + params)
	if err != nil {
		return nil, fmt.Errorf("error while making the request: %w", err)
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return &weather, nil
}

func (ws *WeatherService) getCoordinates(cityName string) ([]CoordinatesResponse, error) {
	params := "?q=" + cityName + "&appid=" + ws.API_KEY
	resp, err := ws.client.Get(ws.API_GEO_V1 + params)
	if err != nil {
		return nil, fmt.Errorf("error while making the request: %w", err)
	}
	defer resp.Body.Close()

	var coordinates []CoordinatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&coordinates); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return coordinates, nil
}

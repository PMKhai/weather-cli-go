package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CoordinatesResponse struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

func getWeather(cityName string) {

	coordinates, err := getCoordinates(cityName)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println(coordinates)
}

func getCoordinates(cityName string) ([]CoordinatesResponse, error) {
	apiKey := getEnvOrDefault("API_KEY", "")
	endpointV1 := getEnvOrDefault("API_GEO_V1", "")

	geoResponse, err := makeGetRequest(endpointV1, map[string]string{
		"q":     cityName,
		"appid": apiKey,
	})

	if err != nil {
		return nil, fmt.Errorf("error while making the request: %w", err)
	}

	defer closeResponseBody(geoResponse)

	body, err := io.ReadAll(geoResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response: %w", err)
	}

	var coordinates []CoordinatesResponse
	if err := json.Unmarshal(body, &coordinates); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return coordinates, nil
}

func makeGetRequest(url string, queryParams map[string]string) (*http.Response, error) {
	query := url + "?" + encodeQueryParams(queryParams)
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func encodeQueryParams(queryParams map[string]string) string {
	values := url.Values{}
	for key, value := range queryParams {
		values.Set(key, value)
	}
	return values.Encode()
}

func closeResponseBody(response *http.Response) {
	if err := response.Body.Close(); err != nil {
		log.Printf("Error while closing response body: %v", err)
	}
}

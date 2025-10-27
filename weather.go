package main

import (
	"fmt"
)

type WeatherResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Current   struct {
		Time               string  `json:"time"`
		Temperature2m      float64 `json:"temperature_2m"`
		RelativeHumidity2m float64 `json:"relative_humidity_2m"`
		ApparentTemp       float64 `json:"apparent_temperature"`
		WindSpeed10m       float64 `json:"wind_speed_10m"`
	} `json:"current"`
	Daily struct {
		Time             []string  `json:"time"`
		TempMax          []float64 `json:"temperature_2m_max"`
		TempMin          []float64 `json:"temperature_2m_min"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
		Sunrise          []string  `json:"sunrise"`
		Sunset           []string  `json:"sunset"`
	} `json:"daily"`
}

func fetchWeather(lat, lon float64, days int) (*WeatherResponse, error) {
	if days < 1 {
		days = 3
	}
	if days > 7 {
		days = 7
	}
	u := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%g&longitude=%g&current=temperature_2m,relative_humidity_2m,apparent_temperature,wind_speed_10m&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,sunrise,sunset&forecast_days=%d&timezone=auto", lat, lon, days)
	var w WeatherResponse
	if err := httpGetJSON(u, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

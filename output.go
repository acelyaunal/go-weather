package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Out struct {
	City  string     `json:"city"`
	Now   NowOut     `json:"now"`
	Daily []DailyOut `json:"daily"`
}

type NowOut struct {
	Temperature     float64 `json:"temperature_c"`
	Apparent        float64 `json:"apparent_c"`
	HumidityPercent float64 `json:"humidity_percent"`
	WindMS          float64 `json:"wind_ms"`
	Timezone        string  `json:"timezone"`
}

type DailyOut struct {
	Date    string  `json:"date"`
	MinC    float64 `json:"min_c"`
	MaxC    float64 `json:"max_c"`
	RainMM  float64 `json:"rain_mm"`
	Sunrise string  `json:"sunrise"`
	Sunset  string  `json:"sunset"`
}

func printNow(city string, w *WeatherResponse) {
	fmt.Printf("📍 %s  (lat: %.3f, lon: %.3f, tz: %s)\n", city, w.Latitude, w.Longitude, w.Timezone)
	fmt.Println("— Şu anki durum —")
	fmt.Printf("🌡  Sıcaklık:      %.1f°C (Hissedilen: %.1f°C)\n", w.Current.Temperature2m, w.Current.ApparentTemp)
	fmt.Printf("💧  Nem:           %%%.0f\n", w.Current.RelativeHumidity2m)
	fmt.Printf("💨  Rüzgar:        %.1f m/s\n", w.Current.WindSpeed10m)
}

func printDaily(days int, w *WeatherResponse) {
	if days > len(w.Daily.Time) {
		days = len(w.Daily.Time)
	}
	fmt.Println("\n— Günlük Tahmin —")
	fmt.Println("Tarih        Min/Max (°C)   Yağış (mm)   Gün Doğumu   Gün Batımı")
	for i := 0; i < days; i++ {
		date := w.Daily.Time[i]
		fmt.Printf("%-12s  %5.1f / %-5.1f   %6.1f   %s   %s\n",
			date, w.Daily.TempMin[i], w.Daily.TempMax[i], w.Daily.PrecipitationSum[i],
			trimTime(w.Daily.Sunrise[i]), trimTime(w.Daily.Sunset[i]),
		)
	}
}

func printJSON(pretty string, w *WeatherResponse, days int) {
	out := Out{
		City: pretty,
		Now: NowOut{
			Temperature:     w.Current.Temperature2m,
			Apparent:        w.Current.ApparentTemp,
			HumidityPercent: w.Current.RelativeHumidity2m,
			WindMS:          w.Current.WindSpeed10m,
			Timezone:        w.Timezone,
		},
	}
	limit := days
	if limit > len(w.Daily.Time) {
		limit = len(w.Daily.Time)
	}
	for i := 0; i < limit; i++ {
		out.Daily = append(out.Daily, DailyOut{
			Date:    w.Daily.Time[i],
			MinC:    w.Daily.TempMin[i],
			MaxC:    w.Daily.TempMax[i],
			RainMM:  w.Daily.PrecipitationSum[i],
			Sunrise: trimTime(w.Daily.Sunrise[i]),
			Sunset:  trimTime(w.Daily.Sunset[i]),
		})
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(out)
}

func trimTime(s string) string {
	if idx := strings.Index(s, "T"); idx >= 0 && idx+1 < len(s) {
		return s[idx+1:]
	}
	return s
}

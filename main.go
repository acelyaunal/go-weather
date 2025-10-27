package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	city, days, asJSON := supportArgs()
	if strings.TrimSpace(city) == "" {
		fmt.Println("Kullanım:")
		fmt.Println("  go run . -days 5 \"Berlin, DE\"")
		fmt.Println("  go run . \"Istanbul\" -days 3")
		fmt.Println("  go run . -json \"Ankara\"")
		os.Exit(2)
	}
	lat, lon, pretty, _, err := geocodeCity(city)
	if err != nil {
		fmt.Println("Hata (geocode):", err)
		os.Exit(1)
	}
	w, err := fetchWeather(lat, lon, days)
	if err != nil {
		fmt.Println("Hata (weather):", err)
		os.Exit(1)
	}
	if asJSON {
		printJSON(pretty, w, days)
		return
	}
	printNow(pretty, w)
	printDaily(days, w)
}

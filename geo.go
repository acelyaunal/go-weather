package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
		Admin1    string  `json:"admin1"`
		Timezone  string  `json:"timezone"`
	} `json:"results"`
}

func httpGetJSON(u string, v any) error {
	c := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest(http.MethodGet, u, nil)
	req.Header.Set("User-Agent", "go-weather-cli/1.0")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("HTTP %d", res.StatusCode)
	}
	return json.NewDecoder(res.Body).Decode(v)
}

func splitCityCountry(input string) (city, wantCountry string) {
	s := strings.TrimSpace(input)
	if i := strings.Index(s, ","); i >= 0 {
		city = strings.TrimSpace(s[:i])
		wantCountry = strings.TrimSpace(s[i+1:])
		return
	}
	parts := strings.Fields(s)
	if len(parts) >= 2 {
		city = parts[0]
		wantCountry = parts[len(parts)-1]
		return
	}
	city = s
	return
}

func geocodeCity(name string) (lat, lon float64, niceName, tz string, err error) {
	if strings.TrimSpace(name) == "" {
		return 0, 0, "", "", errors.New("şehir adı boş olamaz")
	}

	cityOnly, wantCountry := splitCityCountry(name)

	u := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=5&language=tr&format=json",
		url.QueryEscape(cityOnly))
	var geo GeoResponse
	if err = httpGetJSON(u, &geo); err != nil {
		return
	}
	if len(geo.Results) == 0 {
		return 0, 0, "", "", fmt.Errorf("şehir bulunamadı: %s", name)
	}

	best := geo.Results[0]
	if wantCountry != "" {
		for _, r := range geo.Results {
			if countryMatch(r.Country, wantCountry) {
				best = r
				break
			}
		}
	} else {
		for _, r := range geo.Results {
			if r.Country == "Türkiye" {
				best = r
				break
			}
		}
	}

	nice := best.Name
	if best.Admin1 != "" {
		nice += ", " + best.Admin1
	}
	if best.Country != "" {
		nice += ", " + best.Country
	}
	return best.Latitude, best.Longitude, nice, best.Timezone, nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/champ-isaac/weather-service/internal/weather"
	"github.com/champ-isaac/weather-service/pkg/cache"
)

var (
	inMemory    *cache.Cache[*weather.ForecastResponse]
	client      *weather.Client
	forecastFmt = "temperature(°C): %.1f, today forecast: %s, characterization: %s"
)

func main() {
	ctx := context.Background()
	// start cache
	inMemory = cache.New[*weather.ForecastResponse](ctx, 2*time.Minute)
	// start weather client
	client = weather.New()

	http.HandleFunc("/forecast", handleForecast)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleForecast(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	lat := q.Get("latitude")
	lng := q.Get("longitude")
	units := q.Get("units")

	if len(lat) == 0 || len(lng) == 0 {
		http.Error(w, "Missing required parameters.", http.StatusBadRequest)
		return
	}
	if len(units) == 0 {
		units = "si"
	}
	key := fmt.Sprintf("%s,%s,%s", lat, lng, units)
	resp, ok := inMemory.Get(key)
	if !ok {
		log.Println("cache miss")
		var err error
		resp, err = client.Forecast(lat, lng, units)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inMemory.Set(key, resp)
	} else {
		log.Println("cache hit")
	}
	// cache hit or fresh fetched response
	retBs := convertTemperature(resp)

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(retBs); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func convertTemperature(forecast *weather.ForecastResponse) []byte {
	firstPeriod := forecast.Periods[0]
	temperature := firstPeriod.Temperature.Value
	var characterization string
	//right now only implement units = si
	if temperature < 5 {
		characterization = "extremely cold"
	} else if temperature >= 5 && temperature < 10 {
		characterization = "cold"
	} else if temperature >= 10 && temperature < 26 {
		characterization = "moderate"
	} else if temperature >= 26 && temperature < 40 {
		characterization = "hot"
	} else {
		characterization = "extremely hot"
	}
	f := fmt.Sprintf(forecastFmt, temperature, firstPeriod.ShortForecast, characterization)
	return []byte(f)
}

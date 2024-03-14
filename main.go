package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kobie/tracker/models"
	"github.com/kobie/tracker/utils"
)

// Decoder: https://pkg.go.dev/encoding/json#Decoder
func loadEnv(fileName string) (*models.ApiStruct, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	var apiKey models.ApiStruct
	err = decoder.Decode(&apiKey)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &apiKey, nil
}

func convertCityToCoordinate(city string) (float64, float64, error) {
	var lat, long float64
	apiKey, err := loadEnv("api_config")
	if err != nil {
		return lat, long, err
	}

	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=" + city + "&limit=1&appid=" + apiKey.OpenWeatherMapKey)
	if err != nil {
		return lat, long, err
	}
	defer resp.Body.Close()

	var coordinates []models.CityInfo
	err = json.NewDecoder(resp.Body).Decode(&coordinates)

	if err != nil {
		return lat, long, err
	}
	cityCoordinate := coordinates[0]

	return cityCoordinate.Lat, cityCoordinate.Lon, nil
}

func queryForecastForUpcomingDays(city string) (*models.ForecastCollection, error) {
	lat, long, err := convertCityToCoordinate(city)
	if err != nil {
		return nil, err
	}

	apiKey, err := loadEnv("api_config")
	if err != nil {
		return nil, err
	}

	forecastCollection := make(map[string]*models.DailyForecast)
	day := time.Now()
	for i := 0; i < 3; i++ {
		url := fmt.Sprintf(
			"https://api.openweathermap.org/data/3.0/onecall/day_summary?lat=%v&lon=%v&date=%v&appid=%v",
			lat,
			long,
			day.Format("2006-01-02"),
			apiKey.OpenWeatherMapKey,
		)

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var forecastData models.DailyForecast
		err = json.NewDecoder(resp.Body).Decode(&forecastData)
		if err != nil {
			return nil, err
		}

		forecastCollection[day.Weekday().String()] = &forecastData
		day = day.AddDate(0, 0, 1)
	}



	return utils.ConvertForecastData(forecastCollection), nil
}

func queryWeatherByCity(city string) (*models.WeatherData, error) {
	apiKey, err := loadEnv("api_config")
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey.OpenWeatherMapKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data models.WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := queryWeatherByCity(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/forecast/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := queryForecastForUpcomingDays(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

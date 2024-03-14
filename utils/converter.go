package utils

import "github.com/kobie/tracker/models"

func ConvertForecastData(weatherApiForecast map[string]*models.DailyForecast) *models.ForecastCollection {
	totalHumidity, totalMaxTemp, totalMinTemp := 0.0,0,0

	forecastCollection := models.ForecastCollection{DailyForecast: make(map[string]*models.Forecast)}
	for day, apiForecast := range weatherApiForecast {
		temp := getTemperature(
			apiForecast.Tempature.Min,
			apiForecast.Tempature.Max,
			apiForecast.Tempature.Afternoon,
			apiForecast.Tempature.Night,
			apiForecast.Tempature.Morning,
		)
		internalForecast := models.Forecast{
			Temperature: temp,
		}

		internalForecast.HumidityPrecentage = apiForecast.Humidity.Afternoon
		internalForecast.PercipitationPercentage = apiForecast.Precipitation.Total
		forecastCollection.DailyForecast[day] = &internalForecast

		totalHumidity += apiForecast.Humidity.Afternoon
		totalMaxTemp += temp.Max
		totalMinTemp += temp.Min
	}

	amountOfDays := len(weatherApiForecast)
	forecastCollection.AverageHumidity = int(totalHumidity)/amountOfDays
	forecastCollection.AverageTempMax = totalMaxTemp/amountOfDays
	forecastCollection.AverageTempMin = totalMinTemp/amountOfDays

	return &forecastCollection
}

func getTemperature(min, max, afternoon, night, morning float64) *models.Temperature {
	return &models.Temperature{
		Min: convertKelvinToFahrenheit(min),
		Max: convertKelvinToFahrenheit(max),
		Morning: convertKelvinToFahrenheit(morning),
		Afternoon: convertKelvinToFahrenheit(morning),
		Night: convertKelvinToFahrenheit(night),
	}
}

func convertKelvinToFahrenheit(kelvin float64) int {
	celsius := kelvin - 273.15
	fahrenheit := celsius * 1.80 + 32.00
	return int(fahrenheit)
}
package models

type ApiStruct struct {
	OpenWeatherMapKey string `json:"OpenWeatherMapApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type DailyForecast struct {
	Date string `json:"date"`
	Humidity struct {
		Afternoon float64 `json:"afternoon"`
	} `json:"humidity"`

	Precipitation struct {
		Total float64 `json:"total"`
	} `json:"precipitation"`

	Tempature struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
		Afternoon float64 `json:"afternoon"` 
		Night float64 `json:"night"`
		Morning float64 `json:"morning"`
	} `json:"temperature"`
}

type CityInfo struct {
	Name string `json:"name"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"long"`
}
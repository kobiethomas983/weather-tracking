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

type CurrentForecast struct {
	Timezone string `json:"timezone"`
	CurrentOverCast CurrentOverCast `json:"current"`
}

type CurrentOverCast struct {
	Temp float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Clouds float64 `json:"clouds"`
	WindSpeed float64 `json:"wind_speed"`
	Weather []CurrentWeather `json:"weather"`
}

type CurrentWeather struct {
	Main string `json:"main"`
	Description string `json:"description"`
}
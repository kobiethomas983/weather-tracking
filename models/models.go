package models

type Forecast struct {
	HumidityPrecentage float64 `json:"humidityPercentage"`
	PercipitationPercentage float64 `json:"percipitationPercentage"`
	Temperature *Temperature `json:"temperature"`
}

type Temperature struct {
	Min int
	Max int
	Morning int
	Afternoon int
	Night int
}


type ForecastCollection struct {
	DailyForecast map[string]*Forecast `json:"upcomingForecast"`
	AverageHumidity int `json:"averageHumidityPercentage"`
	AverageTempMax int `json:"averageTemperatureHigh"`
	AverageTempMin int `json:"averahgeTemperatureMinimum"`
}
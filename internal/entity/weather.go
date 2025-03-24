package entity

type TemperatureUnit string

type Weather struct {
	Temperature float64 `json:"temperature"`
}

func (w *Weather) NewWeather(temp float64) *Weather {
	return &Weather{
		Temperature: temp,
	}
}

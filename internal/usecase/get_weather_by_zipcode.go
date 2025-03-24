package usecase

import (
	"log"

	"github.com/egon89/go-zipcode-weather/internal/ports"
)

type GetWeatherByZipcode struct {
	LocationPort ports.LocationPort
}

type GetWeatherByZipcodeOutputDto struct {
	TempCelcius    float64
	TempFahrenheit float64
	TempKelvin     float64
}

type GetWeatherByZipcodeInterface interface {
	Execute(zipcode string) (GetWeatherByZipcodeOutputDto, error)
}

func NewGetWeatherByZipcode(locationPort ports.LocationPort) *GetWeatherByZipcode {
	return &GetWeatherByZipcode{
		LocationPort: locationPort,
	}
}

func (g *GetWeatherByZipcode) Execute(zipcode string) (GetWeatherByZipcodeOutputDto, error) {
	city, err := g.LocationPort.GetCityNameByZipcode(zipcode)
	if err != nil {
		return GetWeatherByZipcodeOutputDto{}, err
	}
	log.Printf("getting weather for city %s\n", city)

	return GetWeatherByZipcodeOutputDto{
		TempCelcius:    0,
		TempFahrenheit: 0,
		TempKelvin:     0,
	}, nil
}

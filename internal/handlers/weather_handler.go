package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/go-zipcode-weather/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type WeatherHandler struct {
	usecase usecase.GetWeatherByZipcodeInterface
}

func NewWeatherHandler(getWeatherByZipcode usecase.GetWeatherByZipcodeInterface) *WeatherHandler {
	return &WeatherHandler{
		usecase: getWeatherByZipcode,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipcodeStr := chi.URLParam(r, "zipcode")

	output, err := h.usecase.Execute(zipcodeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(output)
}

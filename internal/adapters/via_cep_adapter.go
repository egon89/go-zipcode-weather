package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/egon89/go-zipcode-weather/internal/config"
)

type ViaCepAdapter struct{}

type viaCepResponse struct {
	Localidade string `json:"localidade"`
}

func NewViaCepAdapter() *ViaCepAdapter {
	return &ViaCepAdapter{}
}

func (vc *ViaCepAdapter) GetCityNameByZipcode(zipcode string) (string, error) {
	log.Printf("[viaCep] getting city name for zipcode %s\n", zipcode)

	url := fmt.Sprintf("%s/%s/json", config.ViaCepBaseURL, zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response viaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if response.Localidade == "" {
		return "", fmt.Errorf("city not found for zipcode %s", zipcode)
	}

	return response.Localidade, nil
}

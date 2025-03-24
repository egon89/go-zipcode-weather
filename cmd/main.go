package main

import (
	"log"
	"net/http"

	"github.com/egon89/go-zipcode-weather/internal/adapters"
	"github.com/egon89/go-zipcode-weather/internal/config"
	"github.com/egon89/go-zipcode-weather/internal/handlers"
	"github.com/egon89/go-zipcode-weather/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()
	port := config.GetEnv("PORT", "8080")

	viaCepAdapter := adapters.NewViaCepAdapter()
	weatherHandler := handlers.NewWeatherHandler(
		usecase.NewGetWeatherByZipcode(viaCepAdapter))

	r := chi.NewRouter()

	r.Get("/weather/{zipcode}", weatherHandler.GetWeather)

	r.Get("/products", handlers.GetProducts)
	r.Get("/products/{id}", handlers.GetProduct)
	r.Post("/products", handlers.CreateProduct)
	r.Put("/products/{id}", handlers.UpdateProduct)
	r.Delete("/products/{id}", handlers.DeleteProduct)

	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

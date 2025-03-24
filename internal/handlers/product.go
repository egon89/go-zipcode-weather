package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/egon89/go-zipcode-weather/internal/entity"
	"github.com/go-chi/chi/v5"
)

var products = []entity.Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	for _, p := range products {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct entity.Product
	json.NewDecoder(r.Body).Decode(&newProduct)
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	json.NewEncoder(w).Encode(newProduct)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var updatedProduct entity.Product
	json.NewDecoder(r.Body).Decode(&updatedProduct)
	updatedProduct.ID = id

	for i, p := range products {
		if p.ID == id {
			products[i] = updatedProduct
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

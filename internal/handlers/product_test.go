package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
    req, _ := http.NewRequest("GET", "/products", nil)
    res := httptest.NewRecorder()

    GetProducts(res, req)

    assert.Equal(t, http.StatusOK, res.Code)
    assert.Contains(t, res.Body.String(), "Laptop")
}

func TestCreateProduct(t *testing.T) {
    body := strings.NewReader(`{"name":"Phone","price":699.99}`)
    req, _ := http.NewRequest("POST", "/products", body)
    req.Header.Set("Content-Type", "application/json")
    res := httptest.NewRecorder()

    CreateProduct(res, req)

    assert.Equal(t, http.StatusOK, res.Code)
    assert.Contains(t, res.Body.String(), "Phone")
}

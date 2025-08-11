package product

import (
	"bytes"
	"encoding/json"
	"go-marketplace/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {

	repo := NewProductRepository()
	handler := NewHandler(repo)

	req, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err, "Failed to create request")
	rr := httptest.NewRecorder()

	handler.ListProductsHandler(rr, req)

	// Check status OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200 Ok")

	products := repo.GetAll()
	expected, err := json.Marshal(products)
	assert.NoError(t, err, "Failed to marshal expected products")

	assert.JSONEq(t, string(expected), rr.Body.String(), "Response body does not match expected products")

}

func TestCreateProducts(t *testing.T) {
	repo := NewProductRepository()
	handler := NewHandler(repo)

	newProduct := models.Product{
		Name:        "Keyboard",
		Price:       20000,
		Description: "Mechanical keyboard",
	}

	productJSON, err := json.Marshal(newProduct)
	assert.NoError(t, err, "Failed to marshal new product")

	body := bytes.NewBuffer(productJSON)
	req, err := http.NewRequest("POST", "/products", body)
	assert.NoError(t, err, "Failed to create request")
	rr := httptest.NewRecorder()

	handler.CreateProductHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201 Created")
	assert.NotEmpty(t, rr.Body.String(), "Response body should not be empty")
}

func TestGetProductHandler(t *testing.T) {
	repo := NewProductRepository()
	handler := NewHandler(repo)

	req, err := http.NewRequest("GET", "/products/1", nil)
	assert.NoError(t, err, "Failed to create request")

	rr := httptest.NewRecorder()

	handler.GetProductHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200 Ok")

	product, err := repo.GetByID(1)
	assert.NoError(t, err, "Product should exist for test")

	expectedJSON, err := json.Marshal(product)
	assert.NoError(t, err, "Failed to marshal expected product")

	assert.JSONEq(t, string(expectedJSON), rr.Body.String(), "Response body does not match expected product")
}

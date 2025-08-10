package product

import (
	"encoding/json"
	"go-marketplace/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	products := []models.Product{
		{ID: 1, Name: "Notebook", Price: 500000, Description: "Notebook gamer"},
		{ID: 2, Name: "Mouse", Price: 15000, Description: "Notebook gamer"},
	}

	expected, err := json.Marshal(products)
	assert.NoError(t, err, "Failed to marshal expected products")

	req, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err, "Failed to create request")

	rr := httptest.NewRecorder()

	ListProductsHandler(rr, req)

	// Check status OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200 Ok")

	assert.JSONEq(t, string(expected), rr.Body.String(), "Response body does not match expected products")

}

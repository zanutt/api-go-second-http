package product

import (
	"encoding/json"
	"go-marketplace/internal/models"
	"net/http"
)

type Handler struct {
	repo *ProductRepository
}

func NewHandler(repo *ProductRepository) *Handler {
	return &Handler{repo: repo}
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {

	// Set content-type header
	w.Header().Set("Content-Type", "application/json")

	products := []models.Product{}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

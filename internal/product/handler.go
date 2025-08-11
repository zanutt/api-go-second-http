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

func (h *Handler) ListProductsHandler(w http.ResponseWriter, r *http.Request) {

	// Set content-type header
	w.Header().Set("Content-Type", "application/json")

	products := h.repo.GetAll()

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdProduct := h.repo.AddProduct(newProduct)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdProduct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

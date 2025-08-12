package product

import (
	"encoding/json"
	"go-marketplace/internal/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	if len(pathSegments) < 3 || pathSegments[2] == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	idStr := pathSegments[2]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.repo.GetByID((uint)(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *Handler) UpdateProductHandler(w *httptest.ResponseRecorder, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 || pathSegments[2] == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	idStr := pathSegments[2]
	idUint64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	id := uint(idUint64)

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.repo.UpdateProduct(uint(id), updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

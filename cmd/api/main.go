package main

import (
	"go-marketplace/internal/product"
	"log"
	"net/http"
)

func main() {
	repo := product.NewProductRepository()
	handler := product.NewHandler(repo)

	mux := http.NewServeMux()

	// List all prod
	mux.HandleFunc("GET /products", handler.ListProductsHandler)

	// Create prod
	mux.HandleFunc("POST /products", handler.CreateProductHandler)

	// Get prod by ID
	mux.HandleFunc("GET /products/{id}", handler.GetProductHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

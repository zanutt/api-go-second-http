package main

import (
	"database/sql"
	"fmt"
	"go-marketplace/internal/infrastructure/database"
	"go-marketplace/internal/product"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	dbUser := "postgres"
	dbPass := "1234567"
	dbName := "go_marketplace"
	dbPort := "5433"

	dsn := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}
	log.Println("Successfully connected to the database!")

	repo := database.NewPostgresRepository(db)

	handler := product.NewHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", handler.ListProductsHandler)
	mux.HandleFunc("POST /products", handler.CreateProductHandler)
	mux.HandleFunc("GET /products/{id}", handler.GetProductHandler)
	mux.HandleFunc("PUT /products/{id}", handler.UpdateProductHandler)
	mux.HandleFunc("DELETE /products/{id}", handler.DeleteProductHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

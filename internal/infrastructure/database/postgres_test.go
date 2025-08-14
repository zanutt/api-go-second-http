package database_test

import (
	"database/sql"
	"fmt"
	"go-marketplace/internal/infrastructure/database"
	"go-marketplace/internal/models"
	"log"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
)

const dsn = "host=localhost port=5433 user=postgres password=1234567 dbname=go_marketplace sslmode=disable"

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	assert.NoError(t, err, "Failed to connect to the database")

	fmt.Println("Successfully connected to the database")
}

func TestPostgresRepository_AddProduct(t *testing.T) {
	db, err := sql.Open("postgres", dsn)
	assert.NoError(t, err, "Failed to open database connection")
	defer db.Close()

	_, err = db.Exec("DELETE FROM products")
	assert.NoError(t, err, "Failed to delete products")

	repo := database.NewPostgresRepository(db)

	newProduct := models.Product{
		Name:        "Monitor Gamer",
		Price:       250000,
		Description: "Monitor gamer 144hz",
	}

	addedProduct, err := repo.AddProduct(newProduct)

	assert.NoError(t, err, "Failed to add product")
	assert.NotZero(t, addedProduct.ID, "Product must have a ID")
	assert.Equal(t, newProduct.Name, addedProduct.Name, "Product name should match")
	assert.Equal(t, newProduct.Price, addedProduct.Price, "Product price should match")
}

func TestPostgresRepository_GetAll(t *testing.T) {
	db, err := sql.Open("postgres", dsn)
	assert.NoError(t, err, "Failed to open database connection")
	defer db.Close()

	_, err = db.Exec("DELETE FROM products")
	assert.NoError(t, err, "Failed to delete products")

	repo := database.NewPostgresRepository(db)

	product1 := models.Product{Name: "Produto A", Price: 10000, Description: "Descrição A"}
	product2 := models.Product{Name: "Produto B", Price: 20000, Description: "Descrição B"}
	repo.AddProduct(product1)
	repo.AddProduct(product2)

	allProducts, err := repo.GetAll()
	assert.NoError(t, err, "Failed to get all products")
	assert.Len(t, allProducts, 2, "Should return two products")
}

func TestPostgresRepository_GetByID(t *testing.T) {
	db, err := sql.Open("postgres", dsn)
	assert.NoError(t, err, "Failed to open database connection")
	defer db.Close()

	_, err = db.Exec("DELETE FROM products")
	assert.NoError(t, err, "Failed to delete products")

	repo := database.NewPostgresRepository(db)

	productToFind := models.Product{
		Name:        "Right Product",
		Price:       15000,
		Description: "Description of the right product",
	}

	addedProduct, err := repo.AddProduct(productToFind)
	assert.NoError(t, err, "Failed to add product")

	foundProduct, err := repo.GetByID(addedProduct.ID)
	assert.NoError(t, err, "Failed to get product by ID")
	assert.Equal(t, addedProduct.ID, foundProduct.ID, "Product ID should match")
	assert.Equal(t, addedProduct.Name, foundProduct.Name, "Product name should match")
	assert.Equal(t, addedProduct.Price, foundProduct.Price, "Product price should match")
	assert.Equal(t, addedProduct.Description, foundProduct.Description, "Product description should match")

	_, err = repo.GetByID(9999) // Non-existent ID
	assert.Error(t, err, "Expected error for non-existent product")
	assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows for non-existent product")
}

func TestPostgresRepository_UpdateProduct(t *testing.T) {
	db, err := sql.Open("postgres", dsn)
	assert.NoError(t, err, "Failed to open database connection")
	defer db.Close()

	_, err = db.Exec("DELETE FROM products")
	assert.NoError(t, err, "Failed to delete products")

	repo := database.NewPostgresRepository(db)

	productToUpdate := models.Product{
		Name:  "Old Keyboard",
		Price: 15000,
	}
	addedProduct, err := repo.AddProduct(productToUpdate)
	assert.NoError(t, err, "Failed to add product")

	updatedData := models.Product{
		Name:        "Updated Keyboard",
		Price:       16000,
		Description: "Updated Description",
	}

	updatedProduct, err := repo.UpdateProduct(addedProduct.ID, updatedData)

	assert.NoError(t, err, "Failed to update product")
	assert.Equal(t, updatedData.Name, updatedProduct.Name)
	assert.Equal(t, updatedData.Price, updatedProduct.Price)

	_, err = repo.UpdateProduct(999, updatedData) // Non-existent ID
	assert.Error(t, err, "Expected error for non-existent product")
	assert.EqualError(t, err, "product not found", "Expected 'product not found' for non-existent product")
}

func TestPostgresRepository_DeleteProduct(t *testing.T) {
	db, err := sql.Open("postgres", dsn)
	assert.NoError(t, err, "Failed to open database connection")
	defer db.Close()

	_, err = db.Exec("DELETE FROM products")
	assert.NoError(t, err, "Failed to delete products")

	repo := database.NewPostgresRepository(db)

	productToDelete := models.Product{
		Name:  "NEW Keyboard",
		Price: 15000,
	}
	addedProduct, err := repo.AddProduct(productToDelete)
	assert.NoError(t, err, "Failed to add product")

	err = repo.DeleteProduct(addedProduct.ID)
	assert.NoError(t, err, "Failed to delete product")

	_, err = repo.GetByID(addedProduct.ID)
	assert.Error(t, err, "Expected error for deleted product")
	assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows for deleted product")

	err = repo.DeleteProduct(999) // Non-existent ID
	assert.Error(t, err, "Expected error for non-existent product")
	assert.EqualError(t, err, "product not found", "Expected 'product not found' for non-existent product")

}

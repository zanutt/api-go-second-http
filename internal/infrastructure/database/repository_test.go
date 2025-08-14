package database_test

import (
	"database/sql"
	"go-marketplace/internal/infrastructure/database"
	"go-marketplace/internal/models"
	"go-marketplace/internal/product"
	"os"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setupTestDB(t *testing.T) (product.ProductRepository, func()) {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	// Cria a tabela de produtos
	if _, err := db.Exec(`CREATE TABLE products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`); err != nil {
		db.Close()
		t.Fatalf("Failed to create table: %v", err)
	}

	repo, err := database.NewSQLiteRepository()
	if err != nil {
		db.Close()
		t.Fatalf("Failed to create repository: %v", err)
	}

	cleanup := func() {
		db.Close()
	}

	return repo, cleanup
}

func TestRepository_AddProduct(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

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

func TestRepository_GetAll(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	product1 := models.Product{Name: "Produto A", Price: 10000, Description: "Descrição A"}
	product2 := models.Product{Name: "Produto B", Price: 20000, Description: "Descrição B"}
	repo.AddProduct(product1)
	repo.AddProduct(product2)

	allProducts, err := repo.GetAll()
	assert.NoError(t, err, "Failed to get all products")
	assert.Len(t, allProducts, 2, "Should return two products")
}

func TestRepository_GetByID(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

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

	_, err = repo.GetByID(9999)
	assert.Error(t, err, "Expected error for non-existent product")
	assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows for non-existent product")
}

func TestRepository_UpdateProduct(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

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

	_, err = repo.UpdateProduct(999, updatedData)
	assert.Error(t, err, "Expected error for non-existent product")
	assert.EqualError(t, err, "product not found", "Expected 'product not found' for non-existent product")
}

func TestRepository_DeleteProduct(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

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

	err = repo.DeleteProduct(999)
	assert.Error(t, err, "Expected error for non-existent product")
	assert.EqualError(t, err, "product not found", "Expected 'product not found' for non-existent product")
}

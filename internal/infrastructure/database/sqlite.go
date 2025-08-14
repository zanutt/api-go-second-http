package database

import (
	"database/sql"
	"errors"
	"go-marketplace/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(`CREATE TABLE products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`); err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) AddProduct(product models.Product) (models.Product, error) {
	query := "INSERT INTO products (name, price, description) VALUES (?, ?, ?) RETURNING id"

	result, err := r.db.Exec(query, product.Name, product.Price, product.Description)
	if err != nil {
		return models.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}

	product.ID = uint(id)
	return product, nil
}

func (r *SQLiteRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, description FROM products"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *SQLiteRepository) GetByID(id uint) (models.Product, error) {
	query := "SELECT id, name, price, description FROM products WHERE id = ?"

	var p models.Product

	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *SQLiteRepository) UpdateProduct(id uint, updatedProduct models.Product) (models.Product, error) {
	updatedProduct.ID = id

	query := `UPDATE products SET name = ?, price = ?, description = ? WHERE id = ?`
	result, err := r.db.Exec(query, updatedProduct.Name, updatedProduct.Price, updatedProduct.Description, id)
	if err != nil {
		return models.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}
	if rowsAffected == 0 {
		return models.Product{}, errors.New("product not found")
	}

	return updatedProduct, nil
}

func (r *SQLiteRepository) DeleteProduct(id uint) error {
	query := "DELETE FROM products WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

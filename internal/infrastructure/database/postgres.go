package database

import (
	"database/sql"
	"errors"
	"go-marketplace/internal/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) AddProduct(product models.Product) (models.Product, error) {
	query := "INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id"
	createdProduct := models.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}

	err := r.db.QueryRow(query, product.Name, product.Price, product.Description).Scan(&createdProduct.ID)
	if err != nil {
		return models.Product{}, err
	}
	return createdProduct, nil
}

func (r *PostgresRepository) GetAll() ([]models.Product, error) {
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

func (r *PostgresRepository) GetByID(id uint) (models.Product, error) {
	query := "SELECT id, name, price, description FROM products WHERE id = $1"

	var p models.Product

	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *PostgresRepository) UpdateProduct(id uint, updatedProduct models.Product) (models.Product, error) {
	updatedProduct.ID = id

	query := `UPDATE products SET name = $1, price = $2, description = $3 WHERE id = $4`
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

func (r *PostgresRepository) DeleteProduct(id uint) error {
	query := "DELETE FROM products WHERE id = $1"
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

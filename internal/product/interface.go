package product

import (
	"go-marketplace/internal/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (models.Product, error)
	AddProduct(product models.Product) (models.Product, error)
	UpdateProduct(id uint, product models.Product) (models.Product, error)
	DeleteProduct(id uint) error
}

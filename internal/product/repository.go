package product

import (
	"errors"
	"go-marketplace/internal/models"
)

type InMemoryRepository struct {
	products map[uint]models.Product
	nextID   uint
}

func NewInMemoryRepository() *InMemoryRepository {
	products := map[uint]models.Product{
		1: {ID: 1, Name: "Notebook", Price: 500000, Description: "Notebook gamer"},
		2: {ID: 2, Name: "Mouse", Price: 15000, Description: "Notebook gamer"},
	}

	return &InMemoryRepository{
		products: products,
		nextID:   3,
	}
}

// get all products
func (r *InMemoryRepository) GetAll() ([]models.Product, error) {
	var allProducts []models.Product
	for _, product := range r.products {
		allProducts = append(allProducts, product)
	}
	return allProducts, nil
}

func (r *InMemoryRepository) AddProduct(product models.Product) (models.Product, error) {
	product.ID = r.nextID

	r.products[product.ID] = product

	r.nextID++

	return product, nil
}

func (r *InMemoryRepository) GetByID(id uint) (models.Product, error) {
	product, ok := r.products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (r *InMemoryRepository) UpdateProduct(id uint, updatedProduct models.Product) (models.Product, error) {
	if _, ok := r.products[id]; !ok {
		return models.Product{}, errors.New("product not found")
	}

	updatedProduct.ID = id

	r.products[id] = updatedProduct
	return updatedProduct, nil
}

func (r *InMemoryRepository) DeleteProduct(id uint) error {
	if _, ok := r.products[id]; !ok {
		return errors.New("product not found")
	}

	delete(r.products, id)
	return nil
}

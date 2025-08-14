package product

import (
	"errors"
	"go-marketplace/internal/models"
)

type ProductRepository struct {
	products map[uint]models.Product
	nextID   uint
}

func NewProductRepository() *ProductRepository {
	products := map[uint]models.Product{
		1: {ID: 1, Name: "Notebook", Price: 500000, Description: "Notebook gamer"},
		2: {ID: 2, Name: "Mouse", Price: 15000, Description: "Notebook gamer"},
	}

	return &ProductRepository{
		products: products,
		nextID:   3,
	}
}

// get all products
func (r *ProductRepository) GetAll() []models.Product {
	var allProducts []models.Product
	for _, product := range r.products {
		allProducts = append(allProducts, product)
	}
	return allProducts
}

func (r *ProductRepository) AddProduct(product models.Product) models.Product {
	product.ID = r.nextID

	r.products[product.ID] = product

	r.nextID++

	return product
}

func (r *ProductRepository) GetByID(id uint) (models.Product, error) {
	product, ok := r.products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(id uint, updatedProduct models.Product) (models.Product, error) {
	if _, ok := r.products[id]; !ok {
		return models.Product{}, errors.New("product not found")
	}

	updatedProduct.ID = id

	r.products[id] = updatedProduct
	return updatedProduct, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	if _, ok := r.products[id]; !ok {
		return errors.New("product not found")
	}

	delete(r.products, id)
	return nil
}

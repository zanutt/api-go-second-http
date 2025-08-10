package product

import "go-marketplace/internal/models"

type ProductRepository struct {
	products map[uint]models.Product
}

func NewProductRepository() *ProductRepository {
	products := map[uint]models.Product{
		1: {ID: 1, Name: "Notebook", Price: 500000, Description: "Notebook gamer"},
		2: {ID: 2, Name: "Mouse", Price: 15000, Description: "Notebook gamer"},
	}

	return &ProductRepository{
		products: products,
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

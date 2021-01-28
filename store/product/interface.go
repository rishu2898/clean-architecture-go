package product

import "Project_store/models"

type Store interface {
	GetById(id int64) (models.Product, error)
	InsertProduct(prod models.Product) (int64, error)
	GetByName(ProductName string) (models.Product, error)
}

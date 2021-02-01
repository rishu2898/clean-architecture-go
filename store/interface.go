package store

import "Project_store/models"

type ProductStore interface {
	GetById(id int64) (models.Product, error)
	InsertProduct(prod models.Product) (int64, error)
	GetByName(ProductName string) (models.Product, error)
}


type BrandStore interface {
	GetById(id int64) (models.Brand, error)
	InsertBrand(name string) (int64, error)
	GetByName(brandName string) (models.Brand, error)
}

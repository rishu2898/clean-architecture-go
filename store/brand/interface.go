package brand

import "Project_store/models"

type Store interface {
	GetById(id int64) (models.Brand, error)
	InsertBrand(name string) (int64, error)
	GetByName(brandName string) (models.Brand, error)
}
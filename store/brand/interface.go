package brand

import "Project_store/models"

type Store interface {
	GetById(id int) (models.Brand, error)
	InsertBrand(prod models.Brand) (int64, error)
}
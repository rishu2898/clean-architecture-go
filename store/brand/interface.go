package brand

import "Project_store/models"

type Store interface {
	GetById(id int) (models.Brand, error)
}
package product

import "Project_store/models"

type Store interface {
	GetById(id int) (models.Product, error)
}
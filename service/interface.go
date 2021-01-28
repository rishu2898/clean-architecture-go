package service

import "Project_store/models"

type Service interface {
	 GetProductDetails(id int64) (models.Result, error)
	 InsertProduct(productName string, brandName string) (models.Result, error)
}

package service

import "Project_store/models"

type Service interface {
	 GetProductDetails(id int) models.Result
}

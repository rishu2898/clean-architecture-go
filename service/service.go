package service

import (
	"Project_store/models"
	"Project_store/store/brand"
	"Project_store/store/product"
	"fmt"
)

type Result struct {
	p product.Store
	b brand.Store
}

func New(p product.Store, b brand.Store) Service {
	return &Result{p, b}
}

func (s Result) GetProductDetails(id int) (models.Result, error) {
	var res models.Result
	productResult, err := s.p.GetById(id)
	if err != nil {
		fmt.Println("id not found")
		return models.Result{}, err
	}
	res.Id = productResult.Id
	res.Name = productResult.Name

	brandResult, _ := s.b.GetById(productResult.BrandId)
	res.Bname = brandResult.Name
	fmt.Println(res)
	return res, nil
}
package service

import (
	"Project_store/models"
	"Project_store/store"
	"errors"
)

type Result struct {
	p store.ProductStore
	b store.BrandStore
}

func New(p store.ProductStore, b store.BrandStore) Service {
	return &Result{p, b}
}
func (s Result) InsertProduct(productName string, brandName string) (models.Result, error) {
	brand, err := s.b.GetByName(brandName)
	var bid int64
	bflag, pflag := false, false
	if err != nil {
		bid, _ = s.b.InsertBrand(brandName)
		brand.Id = bid
		brand.Name = brandName
		bflag = true
	}
	product, err := s.p.GetByName(productName)
	if err != nil {
		product.BrandId = bid
		product.Name = productName
		pid, _ := s.p.InsertProduct(product)
		product.Id = pid
		pflag = true
	}
	if pflag == true && bflag == true {
		return models.Result{}, errors.New("product and brand is already present")
	}
	var data models.Result
	data.Id = product.Id
	data.Name = product.Name
	data.Bname = brand.Name
	return data, nil
}
func (s Result) GetProductDetails(id int64) (models.Result, error) {
	var res models.Result
	productResult, err := s.p.GetById(id)
	if err != nil {
		return models.Result{}, err
	}
	res.Id = productResult.Id
	res.Name = productResult.Name

	brandResult, _ := s.b.GetById(productResult.BrandId)
	res.Bname = brandResult.Name
	return res, nil
}

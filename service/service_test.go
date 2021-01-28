package service

import (
	"Project_store/models"
	"Project_store/store/brand"
	"Project_store/store/product"
	"errors"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ps := product.NewMockStore(ctrl)
	bs := brand.NewMockStore(ctrl)
	psr := New(ps, bs)

	product := []models.Product{
		{1, "bat", 1},
		{2, "ball", 2},
		{3, "wicket", 1},
		{},
	}

	brand := []models.Brand{
		{1, "reebok"},
		{2, "sparten"},
		{},
	}

	expect :=[]models.Result {
		{1, "bat", "reebok"},
		{2, "ball", "sparten"},
		{3, "wicket", "reebok"},
		{},
	}
	testcases := []struct {
		id int64
		prod models.Product
		br models.Brand
		expected models.Result
		err error
	} {
		{1, product[0], brand[0], expect[0], nil},
		{2, product[1], brand[1], expect[1], nil},
		{3, product[2], brand[0], expect[2], nil},
		{0, product[3], brand[2], expect[3], errors.New("invalid id")},
	}

	for _, tc := range testcases {
		ps.EXPECT().GetById(tc.id).Return(tc.prod, tc.err)
		if tc.err == nil {
			bs.EXPECT().GetById(tc.prod.BrandId).Return(tc.br, tc.err)
		}
		ans, err := psr.GetProductDetails(tc.id)
		if err != tc.err {
			log.Fatal(err)
		}
		if ans != tc.expected {
			log.Fatal(err)
		}
	}
}

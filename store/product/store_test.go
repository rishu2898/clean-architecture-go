package product

import (
	"Project_store/models"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// function for testing product by id
func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
	}
	defer db.Close()

	brandHandler := New(db)
	testcases := []models.Product {
		{1, "levis", 1},
		{2, "puma", 2},
	}
	for _, tc := range testcases {
		rows := sqlmock.NewRows([]string{"id", "name", "bid"}).AddRow(tc.Id, tc.Name, tc.BrandId)
		query := "SELECT id, name, bid FROM product WHERE id = ?"
		mock.ExpectQuery(query).WithArgs(tc.Id).WillReturnRows(rows)
		res, err := brandHandler.GetById(tc.Id)
		if err != nil {
			log.Fatal(err)
		}
		data := models.Product{tc.Id, tc.Name, tc.BrandId}
		assert.Equal(t, data, res)
	}
}


// function for testing error cases or invalid id
func TestGetByIdErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
	}
	defer db.Close()

	brandHandler := New(db)
	testcases := []struct {
		id  int
		err string
	}{
		{0, "invlaid id"},
	}
	for _, tc := range testcases {
		//rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(tc.Id, tc.Name)
		query := "SELECT id, name, bid FROM product WHERE id = ?"
		mock.ExpectQuery(query).WithArgs(tc.id).WillReturnError(errors.New(tc.err))
		_, err := brandHandler.GetById(tc.id)

		if err.Error() != tc.err {
			log.Fatal(err)
		}
	}
}


// function for testing insert product into product table
func TestInsertProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
	}
	defer db.Close()

	brandHandler := New(db)
	testcases := []models.Product {
		{1, "levis", 1},
		{2, "puma", 2},
	}
	for _, tc := range testcases {
		query := "INSERT INTO product"
		mock.ExpectExec(query).WithArgs(tc.Id, tc.Name, tc.BrandId).WillReturnResult(sqlmock.NewResult(int64(tc.Id), 1))
		res, err := brandHandler.InsertProduct(tc)
		if err != nil || res != int64(tc.Id) {
			log.Fatal(err)
		}
	}
}
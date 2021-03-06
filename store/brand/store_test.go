package brand

import (
	"Project_store/models"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)
// function for testing brand by id
func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
	}
	defer db.Close()

	brandHandler := New(db)
	testcases := []models.Brand{
		{1, "levis"},
		{2, "puma"},
	}
	for _, tc := range testcases {
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(tc.Id, tc.Name)
		query := "SELECT id, name FROM brand WHERE id = ?"
		mock.ExpectQuery(query).WithArgs(tc.Id).WillReturnRows(rows)
		res, err := brandHandler.GetById(tc.Id)
		if err != nil {
			log.Fatal(err)
		}
		data := models.Brand{tc.Id, tc.Name}
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
		id int64
		err string
	} {
		{0, "invlaid id"},
	}
	for _, tc := range testcases {
		//rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(tc.Id, tc.Name)
		query := "SELECT id, name FROM brand WHERE id = ?"
		mock.ExpectQuery(query).WithArgs(tc.id).WillReturnError(errors.New(tc.err))
		_, err := brandHandler.GetById(tc.id)

		if err.Error() != tc.err {
			log.Fatal(err)
		}
	}
}


// function for test insert data into brand table
func TestInsertProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
	}
	defer db.Close()

	brandHandler := New(db)
	testcases := []models.Brand {
		{1, "levis"},
		{2, "puma"},
	}
	for _, tc := range testcases {
		query := "INSERT INTO brand"
		mock.ExpectExec(query).WithArgs(tc.Name).WillReturnResult(sqlmock.NewResult(int64(tc.Id), 1))
		res, err := brandHandler.InsertBrand(tc.Name)
		if err != nil || res != int64(tc.Id) {
			log.Fatal(err)
		}
	}
}


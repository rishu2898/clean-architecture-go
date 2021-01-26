package brand

import (
	"Project_store/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer db.Close()
	prod := New(db)
	testcases := []models.Brand {
		{1, "levis"},
		{2, "wrogn"},
	}
	for _, tc := range testcases {
		rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(tc.Id, tc.Name)
		query := "SELECT id, name FROM brand WHERE id = ?"
		mock.ExpectQuery(query).WithArgs(tc.Id).WillReturnRows(rows)
		prod.GetById(tc.Id)
	}
	prod.GetById(3)
}

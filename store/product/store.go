package product

import (
	"Project_store/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return product{db}
}

func (prod product) GetById(id int) (models.Product, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = prod.db.Query("SELECT id, name, bid, FROM product where id = ?", id)

	if err != nil {
		return models.Product{}, err
	}

	defer rows.Close()

	var products models.Product
	for rows.Next() {
		_ = rows.Scan(&products.Id, &products.Name, &products.BrandId)
	}
	fmt.Println(products)
	return products, nil
}

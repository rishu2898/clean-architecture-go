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

func (prod product) InsertProduct(val models.Product) (int64, error) {
	res, _ := prod.db.Exec("INSERT INTO product(id, name, bid) VALUES(?, ?, ?)", val.Id, val.Name, val.BrandId)
	lid, err := res.LastInsertId()
	if err != nil {
		fmt.Println("record not inserted")
		return -1, err
	}
	return lid, nil
}

func (prod product) GetById(id int) (models.Product, error) {

	rows, err := prod.db.Query("SELECT id, name, bid FROM product WHERE id = ?", id)

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

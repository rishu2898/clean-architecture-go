package brand

import (
	"Project_store/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type brand struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return brand{db}
}


func (prod brand) InsertBrand(val models.Brand) (int64, error) {
	res, _ := prod.db.Exec("INSERT INTO brand(id, name) VALUES(?, ?)", val.Id, val.Name)
	lid, err := res.LastInsertId()
	if err != nil {
		fmt.Println("record not inserted")
		return -1, err
	}
	return lid, nil
}
func (prod brand) GetById(id int) (models.Brand, error) {

	rows, err := prod.db.Query("SELECT id, name FROM brand WHERE id = ?", id)

	if err != nil {
		return models.Brand{}, err
	}

	defer rows.Close()

	var brands models.Brand
	for rows.Next() {
		_ = rows.Scan(&brands.Id, &brands.Name)
	}
	return brands, nil
}


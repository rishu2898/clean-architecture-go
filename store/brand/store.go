package brand
import (
	"Project_store/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type brand struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return brand{db}
}

func (prod brand) GetById(id int) (models.Brand, error) {

	rows, err := prod.db.Query("SELECT id, name, FROM brand where id = ?", id)

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


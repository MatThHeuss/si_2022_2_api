package database

import (
	"database/sql"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"log"
)

type Product struct {
	DB *sql.DB
}

func NewProductDb(db *sql.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) (int64, error) {
	stmt, err := p.DB.Prepare("insert into products (id, name, description, user_id) values (? , ?, ?, ? )")

	if err != nil {
		log.Printf("Error in prepare statement: %s", err)
		return 0, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(product.ID, product.Name, product.Description, product.UserId)
	id, _ := res.LastInsertId()

	if err != nil {
		log.Printf("Error im exec query: %s", err)
		return 0, err
	}

	log.Println("Success in insertion")
	return id, nil

}

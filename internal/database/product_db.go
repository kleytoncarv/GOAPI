package database

import (
	"database/sql"

	"github.com/kleytoncarv/GOAPI/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, price FROM products")
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var products []*entity.Product
    for rows.Next() {
        var product entity.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID); err!= nil {
            return nil, err
        }
        products = append(products, &product)
    }
    return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT id, name, price, category_id FROM products WHERE id = ?", id).Scan




	
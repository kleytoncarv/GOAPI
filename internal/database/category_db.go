package database

import (
	"database/sql"

	"github.com/kleytoncarv/GOAPI/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func newCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error){
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error){
	rows, err := cd.db 
}
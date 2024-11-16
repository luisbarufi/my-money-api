package repository

import "database/sql"

func NewCategoryRepository(dataBase *sql.DB) CategoryRepository {
	return &categoryRepository{
		dataBase,
	}
}

type categoryRepository struct {
	db *sql.DB
}

type CategoryRepository interface {
}

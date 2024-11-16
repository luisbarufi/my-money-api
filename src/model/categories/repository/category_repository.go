package repository

import (
	"database/sql"

	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
)

func NewCategoryRepository(dataBase *sql.DB) CategoryRepository {
	return &categoryRepository{
		dataBase,
	}
}

type categoryRepository struct {
	db *sql.DB
}

type CategoryRepository interface {
	CreateCategoryRepository(
		categoryDomain model.CategoryDomainInterface,
	) (model.CategoryDomainInterface, *rest_err.RestErr)
}

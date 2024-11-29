package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository"
)

func NewCategoryDomainService(
	categoryRepository repository.CategoryRepository,
) CategoryDomainService {
	return &categoryDomainService{categoryRepository}
}

type categoryDomainService struct {
	categoryRepository repository.CategoryRepository
}

type CategoryDomainService interface {
	CreateCategoryService(
		model.CategoryDomainInterface,
	) (model.CategoryDomainInterface, *rest_err.RestErr)

	ListCategoriesByUserIDService(
		userID uint64,
	) ([]model.CategoryDomainInterface, *rest_err.RestErr)

	UpdateCategoryService(uint64, model.CategoryDomainInterface) *rest_err.RestErr

	DeleteCategoryService(uint64) *rest_err.RestErr
}

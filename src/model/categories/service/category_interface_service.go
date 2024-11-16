package service

import "github.com/luisbarufi/my-money-api/src/model/categories/repository"

func NewCategoryDomainService(
	categoryRepository repository.CategoryRepository,
) CategoryDomainService {
	return &categoryDomainService{categoryRepository}
}

type categoryDomainService struct {
	categoryRepository repository.CategoryRepository
}

type CategoryDomainService interface {
}

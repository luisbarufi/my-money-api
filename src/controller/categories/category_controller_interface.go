package controller

import "github.com/luisbarufi/my-money-api/src/model/categories/service"

func NewCategoryControllerInterface(
	categoryServiceInterface service.CategoryDomainService,
) CategoryControllerInterface {
	return &categoryControllerInterface{
		categoryService: categoryServiceInterface,
	}
}

type CategoryControllerInterface interface {
}

type categoryControllerInterface struct {
	categoryService service.CategoryDomainService
}

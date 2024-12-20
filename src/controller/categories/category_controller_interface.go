package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/model/categories/service"
)

func NewCategoryControllerInterface(
	categoryServiceInterface service.CategoryDomainService,
) CategoryControllerInterface {
	return &categoryControllerInterface{
		categoryService: categoryServiceInterface,
	}
}

type CategoryControllerInterface interface {
	CreateCategoryController(c *gin.Context)
	ListCategoriesByUserIDController(c *gin.Context)
	UpdateCategoryController(c *gin.Context)
	DeleteCategoryController(c *gin.Context)
}

type categoryControllerInterface struct {
	categoryService service.CategoryDomainService
}

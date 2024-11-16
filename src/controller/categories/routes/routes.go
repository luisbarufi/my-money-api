package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/categories"
	"github.com/luisbarufi/my-money-api/src/middlewares"
)

func InitRoutes(
	r *gin.RouterGroup,
	categoryController controller.CategoryControllerInterface,
) {
	categoriesGroup := r.Group("/categories")

	{
		categoriesGroup.POST(
			"/",
			middlewares.VerifyTokenMiddleware,
			categoryController.CreateCategoryController,
		)
	}
}

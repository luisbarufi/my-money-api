package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/categories"
)

func InitRoutes(
	r *gin.RouterGroup,
	categoryController controller.CategoryControllerInterface,
) {
}

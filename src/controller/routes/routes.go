package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/users/:id", userController.FindUserByID)
	r.GET("/users/email/:email", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUserController)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUserController)

	r.POST("/login", userController.LoginUser)
}

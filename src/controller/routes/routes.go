package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/controller"
	"github.com/luisbarufi/my-money-api/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/users/:id", model.VerifyTokenMiddleware, userController.FindUserByIDController)
	r.GET("/users/email/:email", model.VerifyTokenMiddleware, userController.FindUserByEmailController)
	r.PUT("/users/:id", model.VerifyTokenMiddleware, userController.UpdateUserController)
	r.DELETE("/users/:id", model.VerifyTokenMiddleware, userController.DeleteUserController)

	r.POST("/users", userController.CreateUserController)
	r.POST("/login", userController.LoginUserController)
}

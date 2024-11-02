package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/users"
	model "github.com/luisbarufi/my-money-api/src/model/users"
)

func InitRoutes(r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	usersGroup := r.Group("/users")
	{
		usersGroup.GET(
			"/:id",
			model.VerifyTokenMiddleware,
			userController.FindUserByIDController,
		)

		usersGroup.GET(
			"/email/:email",
			model.VerifyTokenMiddleware,
			userController.FindUserByEmailController,
		)

		usersGroup.PUT(
			"/:id",
			model.VerifyTokenMiddleware,
			userController.UpdateUserController,
		)

		usersGroup.DELETE(
			"/:id",
			model.VerifyTokenMiddleware,
			userController.DeleteUserController,
		)

		usersGroup.POST("/", userController.CreateUserController)
	}

	r.POST("/login", userController.LoginUserController)

	r.POST("/forgot-password", userController.ForgotPasswordController)

	r.POST("/update-password", userController.UpdatePasswordController)
}

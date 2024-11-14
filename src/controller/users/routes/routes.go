package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/users"
	"github.com/luisbarufi/my-money-api/src/middlewares"
)

func InitRoutes(r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	usersGroup := r.Group("/users")
	{
		usersGroup.GET(
			"/:id",
			middlewares.VerifyTokenMiddleware,
			userController.FindUserByIDController,
		)

		usersGroup.GET(
			"/email/:email",
			middlewares.VerifyTokenMiddleware,
			userController.FindUserByEmailController,
		)

		usersGroup.PUT(
			"/:id",
			middlewares.VerifyTokenMiddleware,
			userController.UpdateUserController,
		)

		usersGroup.DELETE(
			"/:id",
			middlewares.VerifyTokenMiddleware,
			userController.DeleteUserController,
		)

		usersGroup.POST("/", userController.CreateUserController)
	}

	r.POST("/login", userController.LoginUserController)

	r.POST("/forgot-password", userController.ForgotPasswordController)

	r.POST("/update-password", userController.UpdatePasswordController)
}

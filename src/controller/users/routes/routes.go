package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/users"
	model "github.com/luisbarufi/my-money-api/src/model/users"
)

func InitRoutes(r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET(
		"/users/:id",
		model.VerifyTokenMiddleware,
		userController.FindUserByIDController,
	)

	r.GET(
		"/users/email/:email",
		model.VerifyTokenMiddleware,
		userController.FindUserByEmailController,
	)

	r.PUT(
		"/users/:id",
		model.VerifyTokenMiddleware,
		userController.UpdateUserController,
	)

	r.DELETE(
		"/users/:id",
		model.VerifyTokenMiddleware,
		userController.DeleteUserController,
	)

	r.POST("/users", userController.CreateUserController)

	r.POST("/login", userController.LoginUserController)

	r.POST("/forgot-password", userController.ForgotPasswordController)

	r.POST("/update-password", userController.UpdatePasswordController)
}

package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/accounts"
	"github.com/luisbarufi/my-money-api/src/middlewares"
)

func InitRoutes(
	r *gin.RouterGroup,
	accountController controller.AccountControllerInterface,
) {
	accountsGroup := r.Group("/accounts")
	{
		accountsGroup.POST(
			"/",
			middlewares.VerifyTokenMiddleware,
			accountController.CreateAccountController,
		)

		accountsGroup.GET(
			"/user",
			middlewares.VerifyTokenMiddleware,
			accountController.FindAccountsByUserIDController,
		)
	}
}

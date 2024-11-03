package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/accounts"
	userModel "github.com/luisbarufi/my-money-api/src/model/users"
)

func InitRoutes(
	r *gin.RouterGroup,
	accountController controller.AccountControllerInterface,
) {
	accountsGroup := r.Group("/accounts")
	{
		accountsGroup.POST(
			"/",
			userModel.VerifyTokenMiddleware,
			accountController.CreateAccountController,
		)
	}
}

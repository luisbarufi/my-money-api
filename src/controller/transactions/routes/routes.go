package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/transactions"
	"github.com/luisbarufi/my-money-api/src/middlewares"
)

func InitRoutes(
	r *gin.RouterGroup,
	transactionController controller.TransactionControllerInterface,
) {
	transactionGroup := r.Group("/transactions")

	{
		transactionGroup.POST(
			"/",
			middlewares.VerifyTokenMiddleware,
			transactionController.CreateTransactionController,
		)

		transactionGroup.GET(
			"/user",
			middlewares.VerifyTokenMiddleware,
			transactionController.ListTrasactionsByUserIDController,
		)

		transactionGroup.PUT(
			"/:id",
			middlewares.VerifyTokenMiddleware,
			transactionController.UpdateTransactionController,
		)

		transactionGroup.DELETE(
			"/:id",
			middlewares.VerifyTokenMiddleware,
			transactionController.DeleteTransactionController,
		)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/luisbarufi/my-money-api/src/controller/transactions"
)

func InitRoutes(
	r *gin.RouterGroup,
	transactionController controller.TransactionControllerInterface,
) {
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/model/transactions/service"
)

func NewTransactionControllerInterface(
	transactionServiceInterface service.TransactionDomainService,
) TransactionControllerInterface {
	return &transactionControllerInterface{
		transactionService: transactionServiceInterface,
	}
}

type TransactionControllerInterface interface {
	CreateTransactionController(c *gin.Context)
	ListTrasactionsByUserIDController(c *gin.Context)
	UpdateTransactionController(c *gin.Context)
	DeleteTransactionController(c *gin.Context)
}

type transactionControllerInterface struct {
	transactionService service.TransactionDomainService
}

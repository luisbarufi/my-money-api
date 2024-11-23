package controller

import "github.com/luisbarufi/my-money-api/src/model/transactions/service"

func NewTransactionControllerInterface(
	transactionServiceInterface service.TransactionDomainService,
) TransactionControllerInterface {
	return &transactionControllerInterface{
		transactionService: transactionServiceInterface,
	}
}

type TransactionControllerInterface interface {
}

type transactionControllerInterface struct {
	transactionService service.TransactionDomainService
}

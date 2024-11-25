package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository"
)

func NewTransactionDomainService(
	transactionRepository repository.TransactionRepository,
) TransactionDomainService {
	return &transactionDomainService{transactionRepository}
}

type transactionDomainService struct {
	transactionRepository repository.TransactionRepository
}

type TransactionDomainService interface {
	CreateTransactionService(
		model.TransactionDomainInterface,
	) (model.TransactionDomainInterface, *rest_err.RestErr)

	FindTransactionsByUserIDService(
		userID uint64,
	) ([]model.TransactionDomainInterface, *rest_err.RestErr)
}

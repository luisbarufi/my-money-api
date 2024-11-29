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

	ListTransactionsByUserIDService(
		userID uint64,
	) ([]model.TransactionDomainInterface, *rest_err.RestErr)

	UpdateTransactionService(
		transactionID uint64, transactionDomain model.TransactionDomainInterface,
	) *rest_err.RestErr

	DeleteTransactionService(transactionID uint64) *rest_err.RestErr
}

package service

import "github.com/luisbarufi/my-money-api/src/model/transactions/repository"

func NewTransactionDomainService(
	transactionRepository repository.TransactionRepository,
) TransactionDomainService {
	return &transactionDomainService{transactionRepository}
}

type transactionDomainService struct {
	transactionRepository repository.TransactionRepository
}

type TransactionDomainService interface {
}

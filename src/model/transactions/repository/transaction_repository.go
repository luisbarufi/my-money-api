package repository

import (
	"database/sql"

	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
)

func NewTransactionRepository(dataBase *sql.DB) TransactionRepository {
	return &transactionRepository{
		dataBase,
	}
}

type transactionRepository struct {
	db *sql.DB
}

type TransactionRepository interface {
	CreateTransactionRepository(
		transactionDomain model.TransactionDomainInterface,
	) (model.TransactionDomainInterface, *rest_err.RestErr)

	FindTransactionsByUserIDRepository(
		userID uint64,
	) ([]model.TransactionDomainInterface, *rest_err.RestErr)
}

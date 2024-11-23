package repository

import "database/sql"

func NewTransactionRepository(dataBase *sql.DB) TransactionRepository {
	return &transactionRepository{
		dataBase,
	}
}

type transactionRepository struct {
	db *sql.DB
}

type TransactionRepository interface {
}

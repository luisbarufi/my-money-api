package model

import "time"

type TransactionDomainInterface interface {
	GetID() uint64
	SetID(id uint64)
	GetUserID() uint64
	SetUserID(userID uint64)
	GetAccountID() uint64
	SetAccountID(accountID uint64)
	GetCategoryID() uint64
	SetCategoryID(categoryID uint64)
	GetAmount() float64
	SetAmount(amount float64)
	GetDescription() string
	SetDescription(description string)
	GetTransactionType() string
	SetTransactionType(transactionType string)
	GetCreatedAt() time.Time
	SetCreatedAt(created_at time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updated_at time.Time)
}

func NewTrasactionDomain(
	userID uint64,
	accountID uint64,
	categoryID uint64,
	amount float64,
	description string,
	transactionType string,
) TransactionDomainInterface {
	return &transactionDomain{
		userID:          userID,
		accountID:       accountID,
		categoryID:      categoryID,
		amount:          amount,
		description:     description,
		transactionType: transactionType,
	}
}

func NewUpdateTransactionDomain(
	accountID uint64,
	categoryID uint64,
	amount float64,
	description string,
	transactionType string,
) TransactionDomainInterface {
	return &transactionDomain{
		accountID:       accountID,
		categoryID:      categoryID,
		amount:          amount,
		description:     description,
		transactionType: transactionType,
	}
}

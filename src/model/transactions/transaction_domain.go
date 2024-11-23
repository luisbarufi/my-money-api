package model

import "time"

type transactionDomain struct {
	id              uint64
	userID          uint64
	accountID       uint64
	categoryID      uint64
	amount          float64
	description     string
	transactionType string
	createdAt       time.Time
	updatedAt       time.Time
}

func (td *transactionDomain) GetID() uint64 {
	return td.id
}

func (td *transactionDomain) SetID(id uint64) {
	td.id = id
}

func (td *transactionDomain) GetUserID() uint64 {
	return td.userID
}

func (td *transactionDomain) SetUserID(userID uint64) {
	td.userID = userID
}

func (td *transactionDomain) GetAccountID() uint64 {
	return td.accountID
}

func (td *transactionDomain) SetAccountID(accountID uint64) {
	td.accountID = accountID
}

func (td *transactionDomain) GetCategoryID() uint64 {
	return td.categoryID
}

func (td *transactionDomain) SetCategoryID(categoryID uint64) {
	td.accountID = categoryID
}

func (td *transactionDomain) GetAmount() float64 {
	return td.amount
}

func (td *transactionDomain) SetAmount(amount float64) {
	td.amount = amount
}

func (td *transactionDomain) GetDescription() string {
	return td.description
}

func (td *transactionDomain) SetDescription(description string) {
	td.description = description
}

func (td *transactionDomain) GetTransactionType() string {
	return td.transactionType
}

func (td *transactionDomain) SetTransactionType(transactionType string) {
	td.description = transactionType
}

func (td *transactionDomain) GetCreatedAt() time.Time {
	return td.createdAt
}

func (td *transactionDomain) SetCreatedAt(created_at time.Time) {
	td.createdAt = created_at
}

func (td *transactionDomain) GetUpdatedAt() time.Time {
	return td.updatedAt
}

func (td *transactionDomain) SetUpdatedAt(updated_at time.Time) {
	td.updatedAt = updated_at
}

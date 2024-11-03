package model

import "time"

type AccountDomainInterface interface {
	GetID() uint64

	SetID(id uint64)

	GetUserID() uint64

	SetUserID(userID uint64)

	GetAccountName() string

	SetAccountName(accountName string)

	GetBalance() float64

	SetBalance(balance float64)

	GetCreatedAt() time.Time

	SetCreatedAt(created_at time.Time)

	GetUpdatedAt() time.Time

	SetUpdatedAt(updated_at time.Time)
}

func NewAccountDomain(
	userID uint64,
	accountName string,
	balance float64,
) AccountDomainInterface {
	return &accountDomain{
		userID:      userID,
		accountName: accountName,
		balance:     balance,
	}
}

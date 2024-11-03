package model

import "time"

type accountDomain struct {
	id          uint64
	userID      uint64
	accountName string
	balance     float64
	createdAt   time.Time
	updatedAt   time.Time
}

func (ad *accountDomain) GetID() uint64 {
	return ad.id
}

func (ad *accountDomain) SetID(id uint64) {
	ad.id = id
}

func (ad *accountDomain) GetUserID() uint64 {
	return ad.userID
}

func (ad *accountDomain) SetUserID(userID uint64) {
	ad.userID = userID
}

func (ad *accountDomain) GetAccountName() string {
	return ad.accountName
}

func (ad *accountDomain) SetAccountName(accountName string) {
	ad.accountName = accountName
}

func (ad *accountDomain) GetBalance() float64 {
	return ad.balance
}

func (ad *accountDomain) SetBalance(balance float64) {
	ad.balance = balance
}

func (ad *accountDomain) GetCreatedAt() time.Time {
	return ad.createdAt
}

func (ad *accountDomain) SetCreatedAt(created_at time.Time) {
	ad.createdAt = created_at
}

func (ad *accountDomain) GetUpdatedAt() time.Time {
	return ad.updatedAt
}

func (ad *accountDomain) SetUpdatedAt(updated_at time.Time) {
	ad.updatedAt = updated_at
}

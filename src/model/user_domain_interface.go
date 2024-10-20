package model

import "time"

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string

	GetID() uint64
	SetID(uint64)
	GetCreatedAt() time.Time
	SetCreatedAt(created_at time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updated_at time.Time)

	EncryptPassword()
}

func NewUserDomain(
	name, email, password string,
) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
	}
}

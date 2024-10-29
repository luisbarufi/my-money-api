package model

import (
	"time"

	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
)

type UserDomainInterface interface {
	GetName() string
	GetNick() string
	GetEmail() string
	GetPassword() string

	GetID() uint64
	SetID(uint64)
	GetCreatedAt() time.Time
	SetCreatedAt(created_at time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updated_at time.Time)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
	GenerateResetToken() (string, *rest_err.RestErr)
}

func NewUserDomain(name, nick, email, password string) UserDomainInterface {
	return &userDomain{
		name:     name,
		nick:     nick,
		email:    email,
		password: password,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserResetPasswordDomain(email string) UserDomainInterface {
	return &userDomain{
		email: email,
	}
}

func NewUserUpdateDomain(name, nick string) UserDomainInterface {
	return &userDomain{
		name: name,
		nick: nick,
	}
}

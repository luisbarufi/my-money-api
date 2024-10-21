package model

import "time"

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
}

func NewUserDomain(name, nick, email, password string) UserDomainInterface {
	return &userDomain{
		name:     name,
		nick:     nick,
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(name, nick string) UserDomainInterface {
	return &userDomain{
		name: name,
		nick: nick,
	}
}

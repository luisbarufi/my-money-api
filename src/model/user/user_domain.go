package model

import "time"

type userDomain struct {
	id        uint64
	name      string
	nick      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func (ud *userDomain) GetID() uint64 {
	return ud.id
}

func (ud *userDomain) SetID(id uint64) {
	ud.id = id
}

func (ud *userDomain) GetCreatedAt() time.Time {
	return ud.createdAt
}

func (ud *userDomain) SetCreatedAt(created_at time.Time) {
	ud.createdAt = created_at
}

func (ud *userDomain) GetUpdatedAt() time.Time {
	return ud.updatedAt
}

func (ud *userDomain) SetUpdatedAt(updated_at time.Time) {
	ud.updatedAt = updated_at
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetNick() string {
	return ud.nick
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

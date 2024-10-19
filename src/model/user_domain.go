package model

import "time"

type userDomain struct {
	id         int64
	name       string
	email      string
	password   string
	created_at time.Time
	updated_at time.Time
}

func (ud *userDomain) GetID() int64 {
	return ud.id
}

func (ud *userDomain) SetID(id int64) {
	ud.id = id
}

func (ud *userDomain) GetCreatedAt() time.Time {
	return ud.created_at
}

func (ud *userDomain) SetCreatedAt(created_at time.Time) {
	ud.created_at = created_at
}

func (ud *userDomain) GetUpdatedAt() time.Time {
	return ud.updated_at
}

func (ud *userDomain) SetUpdatedAt(updated_at time.Time) {
	ud.updated_at = updated_at
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

package model

import "time"

type categoryDomain struct {
	id           uint64
	userID       uint64
	categoryName string
	createdAt    time.Time
	updatedAt    time.Time
}

func (cd *categoryDomain) GetID() uint64 {
	return cd.id
}

func (cd *categoryDomain) SetID(id uint64) {
	cd.id = id
}

func (cd *categoryDomain) GetUserID() uint64 {
	return cd.userID
}

func (cd *categoryDomain) SetUserID(userID uint64) {
	cd.userID = userID
}

func (cd *categoryDomain) GetCategoryName() string {
	return cd.categoryName
}

func (cd *categoryDomain) SetCategoryName(categoryName string) {
	cd.categoryName = categoryName
}

func (cd *categoryDomain) GetCreatedAt() time.Time {
	return cd.createdAt
}

func (cd *categoryDomain) SetCreatedAt(created_at time.Time) {
	cd.createdAt = created_at
}

func (cd *categoryDomain) GetUpdatedAt() time.Time {
	return cd.updatedAt
}

func (cd *categoryDomain) SetUpdatedAt(updated_at time.Time) {
	cd.updatedAt = updated_at
}

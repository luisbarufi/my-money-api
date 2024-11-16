package model

import "time"

type CategoryDomainInterface interface {
	GetID() uint64
	SetID(id uint64)
	GetUserID() uint64
	SetUserID(userID uint64)
	GetCategoryName() string
	SetCategoryName(categoryName string)
	GetCreatedAt() time.Time
	SetCreatedAt(created_at time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updated_at time.Time)
}

func NewCategoryDomain(
	userID uint64,
	categoryName string,
) CategoryDomainInterface {
	return &categoryDomain{
		userID:       userID,
		categoryName: categoryName,
	}
}

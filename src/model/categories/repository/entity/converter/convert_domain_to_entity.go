package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity"
)

func ConvertDomainToEntity(domain model.CategoryDomainInterface) *entity.CategoryEntity {
	return &entity.CategoryEntity{
		ID:           domain.GetID(),
		UserID:       domain.GetUserID(),
		CategoryName: domain.GetCategoryName(),
		CreatedAt:    domain.GetCreatedAt(),
		UpdatedAt:    domain.GetUpdatedAt(),
	}
}

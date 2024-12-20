package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity"
)

func ConvertEntityToDomain(entity entity.CategoryEntity) model.CategoryDomainInterface {
	domain := model.NewCategoryDomain(
		entity.UserID,
		entity.CategoryName,
	)
	domain.SetID(entity.ID)
	domain.SetCreatedAt(entity.CreatedAt)
	domain.SetUpdatedAt(entity.UpdatedAt)

	return domain
}

func ConvertEntitiesToDomains(entities []entity.CategoryEntity) []model.CategoryDomainInterface {
	var domains []model.CategoryDomainInterface
	for _, entity := range entities {
		domain := model.NewCategoryDomain(
			entity.UserID,
			entity.CategoryName,
		)
		domain.SetID(entity.ID)
		domain.SetCreatedAt(entity.CreatedAt)
		domain.SetUpdatedAt(entity.UpdatedAt)
		domains = append(domains, domain)
	}
	return domains
}

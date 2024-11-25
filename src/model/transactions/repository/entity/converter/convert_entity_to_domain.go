package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity"
)

func ConvertEntityToDomain(entity entity.TransactionEntity) model.TransactionDomainInterface {
	domain := model.NewTrasactionDomain(
		entity.UserID,
		entity.AccountID,
		entity.CategoryID,
		entity.Amount,
		entity.Description,
		entity.TransactionType,
	)
	domain.SetID(entity.ID)
	domain.SetCreatedAt(entity.CreatedAt)
	domain.SetUpdatedAt(entity.UpdatedAt)

	return domain
}

func ConvertEntitiesToDomains(entities []entity.TransactionEntity) []model.TransactionDomainInterface {
	var domains []model.TransactionDomainInterface
	for _, entity := range entities {
		domain := model.NewTrasactionDomain(
			entity.UserID,
			entity.AccountID,
			entity.CategoryID,
			entity.Amount,
			entity.Description,
			entity.TransactionType,
		)
		domain.SetID(entity.ID)
		domain.SetCreatedAt(entity.CreatedAt)
		domain.SetUpdatedAt(entity.UpdatedAt)
		domains = append(domains, domain)
	}
	return domains
}

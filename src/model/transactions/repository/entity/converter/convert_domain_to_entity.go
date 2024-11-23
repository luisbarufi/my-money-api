package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity"
)

func ConvertDomainToEntity(domain model.TransactionDomainInterface) *entity.TransactionEntity {
	return &entity.TransactionEntity{
		ID:              domain.GetID(),
		UserID:          domain.GetUserID(),
		AccountID:       domain.GetAccountID(),
		CategoryID:      domain.GetCategoryID(),
		Amount:          domain.GetAmount(),
		TransactionType: domain.GetTransactionType(),
		Description:     domain.GetDescription(),
		CreatedAt:       domain.GetCreatedAt(),
		UpdatedAt:       domain.GetUpdatedAt(),
	}
}

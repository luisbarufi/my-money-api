package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity"
)

func ConvertDomainToEntity(domain model.AccountDomainInterface) *entity.AccountEntity {
	return &entity.AccountEntity{
		ID:          domain.GetID(),
		UserID:      domain.GetUserID(),
		AccountName: domain.GetAccountName(),
		Balance:     domain.GetBalance(),
		CreatedAt:   domain.GetCreatedAt(),
		UpdatedAt:   domain.GetUpdatedAt(),
	}
}

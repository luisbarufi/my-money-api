package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity"
)

func ConvertEntityToDomain(entity entity.AccountEntity) model.AccountDomainInterface {
	domain := model.NewAccountDomain(
		entity.UserID,
		entity.AccountName,
		entity.Balance,
	)
	domain.SetID(entity.ID)
	domain.SetCreatedAt(entity.CreatedAt)
	domain.SetUpdatedAt(entity.UpdatedAt)

	return domain
}

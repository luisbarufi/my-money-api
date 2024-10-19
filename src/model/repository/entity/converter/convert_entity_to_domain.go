package converter

import (
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Name,
		entity.Email,
		entity.Password,
	)

	domain.SetID(entity.ID)
	domain.SetCreatedAt(entity.CreatedAt)
	domain.SetUpdatedAt(entity.UpdatedAt)

	return domain
}

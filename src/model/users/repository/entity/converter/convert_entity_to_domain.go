package converter

import (
	model "github.com/luisbarufi/my-money-api/src/model/users"
	"github.com/luisbarufi/my-money-api/src/model/users/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Name,
		entity.Nick,
		entity.Email,
		entity.Password,
	)
	domain.SetID(entity.ID)
	domain.SetCreatedAt(entity.CreatedAt)
	domain.SetUpdatedAt(entity.UpdatedAt)
	return domain
}

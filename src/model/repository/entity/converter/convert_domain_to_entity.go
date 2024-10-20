package converter

import (
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        domain.GetID(),
		Name:      domain.GetName(),
		Nick:      domain.GetNick(),
		Email:     domain.GetEmail(),
		Password:  domain.GetPassword(),
		CreatedAt: domain.GetCreatedAt(),
		UpdatedAt: domain.GetUpdatedAt(),
	}
}

package view

import (
	"github.com/luisbarufi/my-money-api/src/controller/model/response"
	"github.com/luisbarufi/my-money-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.GetID(),
		Name:      userDomain.GetName(),
		Email:     userDomain.GetEmail(),
		CreatedAt: userDomain.GetCreatedAt(),
		UpdatedAt: userDomain.GetUpdatedAt(),
	}
}

package view

import (
	"github.com/luisbarufi/my-money-api/src/controller/users/model/response"
	model "github.com/luisbarufi/my-money-api/src/model/users"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.GetID(),
		Name:      userDomain.GetName(),
		Nick:      userDomain.GetNick(),
		Email:     userDomain.GetEmail(),
		CreatedAt: userDomain.GetCreatedAt(),
		UpdatedAt: userDomain.GetUpdatedAt(),
	}
}

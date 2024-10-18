package view

import (
	"time"

	"github.com/luisbarufi/my-money-api/src/controller/model/response"
	"github.com/luisbarufi/my-money-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:        0,
		Name:      userDomain.GetName(),
		Email:     userDomain.GetEmail(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

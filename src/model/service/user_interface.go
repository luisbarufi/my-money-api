package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(
		model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id uint64,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	LoginUserServices(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(uint64, model.UserDomainInterface) *rest_err.RestErr

	DeleteUser(uint64) *rest_err.RestErr
}

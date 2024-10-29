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
	CreateUserService(
		model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDService(
		id uint64,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailService(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	ForgotPasswordService(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_err.RestErr)

	LoginUserService(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_err.RestErr)

	UpdateUserService(uint64, model.UserDomainInterface) *rest_err.RestErr

	DeleteUserService(uint64) *rest_err.RestErr
}

package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/database/postgres"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
)

func NewUserRepository(dataBase *postgres.Database) UserRepository {
	return &userRepository{
		dataBase,
	}
}

type userRepository struct {
	db *postgres.Database
}

type UserRepository interface {
	CreateUserRepository(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId uint64,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPasswordRepository(
		email, password string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDRepository(id uint64) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUserRepository(userId uint64) *rest_err.RestErr
}

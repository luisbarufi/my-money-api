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
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(id uint64) (model.UserDomainInterface, *rest_err.RestErr)
}

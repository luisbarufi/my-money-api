package repository

import (
	"database/sql"

	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/user"
)

func NewUserRepository(dataBase *sql.DB) UserRepository {
	return &userRepository{
		dataBase,
	}
}

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	CreateUserRepository(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUserRepository(
		userId uint64,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPasswordRepository(
		email, password string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDRepository(id uint64) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUserRepository(userId uint64) *rest_err.RestErr

	UpdatePasswordRepository(
		token string, userDomain model.UserDomainInterface,
	) *rest_err.RestErr
}

package repository

import (
	"database/sql"

	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
)

func NewAccountRepository(dataBase *sql.DB) AccountRepository {
	return &accountRepository{
		dataBase,
	}
}

type accountRepository struct {
	db *sql.DB
}

type AccountRepository interface {
	CreateAccountRepository(
		accountDomain model.AccountDomainInterface,
	) (model.AccountDomainInterface, *rest_err.RestErr)

	FindAccountsByUserIDRepository(
		userID uint64,
	) ([]model.AccountDomainInterface, *rest_err.RestErr)

	UpdateAccountRepository(
		accountId uint64, accountDomain model.AccountDomainInterface,
	) *rest_err.RestErr

	DeleteAccountRepository(accountId uint64) *rest_err.RestErr
}

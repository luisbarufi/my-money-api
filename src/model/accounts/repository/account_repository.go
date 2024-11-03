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
}

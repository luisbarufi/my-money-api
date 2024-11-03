package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository"
)

func NewAccountDomainService(
	accountRepository repository.AccountRepository,
) AccountDomainService {
	return &accountDomainService{accountRepository}
}

type accountDomainService struct {
	accountRepository repository.AccountRepository
}

type AccountDomainService interface {
	CreateAccountService(
		model.AccountDomainInterface,
	) (model.AccountDomainInterface, *rest_err.RestErr)
}
package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"go.uber.org/zap"
)

func (ad *accountDomainService) CreateAccountService(
	accountDomain model.AccountDomainInterface,
) (model.AccountDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateAccountService",
		zap.String("journey", "createAccount"),
	)

	accountDomainRepository, err := ad.accountRepository.CreateAccountRepository(accountDomain)

	if err != nil {
		logger.Error("Error trying to call CreateAccountRepository",
			err,
			zap.String("journey", "createAccount"),
		)

		return nil, err
	}

	logger.Info(
		"CreateAccountService executed successfully",
		zap.String("userId", fmt.Sprintf("%d", accountDomainRepository.GetID())),
		zap.String("journey", "createUser"),
	)

	return accountDomainRepository, nil
}

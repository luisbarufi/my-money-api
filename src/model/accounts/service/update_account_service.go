package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"go.uber.org/zap"
)

func (ad *accountDomainService) UpdateAccountService(
	accountId uint64, accountDomain model.AccountDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateAccountService",
		zap.String("journey", "updateAccount"),
	)

	err := ad.accountRepository.UpdateAccountRepository(accountId, accountDomain)

	if err != nil {
		logger.Error(
			"Error trying to call UpdateAccountRepository",
			err, zap.String("journey", "updateAccount"),
		)

		return err
	}

	logger.Info(
		"UpdateAccountService executed successfully",
		zap.String("journey", "updateAccount"),
	)

	return nil
}

package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ad *accountDomainService) DeleteAccountService(
	accountId uint64,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteAccountService",
		zap.String("journey", "deleteAccount"),
	)

	err := ad.accountRepository.DeleteAccountRepository(accountId)

	if err != nil {
		logger.Error(
			"Error trying to call DeleteAccountRepository",
			err,
			zap.String("journey", "deleteAccount"),
		)

		return err
	}

	logger.Info(
		"DeleteAccountService executed successfully",
		zap.String("accountId", fmt.Sprintf("%d", accountId)),
		zap.String("journey", "deleteAccount"),
	)

	return nil
}

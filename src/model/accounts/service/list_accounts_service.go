package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"go.uber.org/zap"
)

func (ad *accountDomainService) ListAccountsByUserIDService(
	userID uint64,
) ([]model.AccountDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init ListAccountsByUserIDService",
		zap.String("journey", "listAccountsByUserID"),
	)

	logger.Info(
		"ListAccountsByUserIDService executed successfully",
		zap.String("journey", "listAccountsByUserID"),
	)

	return ad.accountRepository.ListAccountsByUserIDRepository(userID)
}

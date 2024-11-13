package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"go.uber.org/zap"
)

func (ad *accountDomainService) FindAccountsByUserIDService(
	userID uint64,
) ([]model.AccountDomainInterface, *rest_err.RestErr) {

	logger.Info(
		"Init FindAccountsByUserIDService",
		zap.String("journey", "findAccountsByuserID"),
	)

	logger.Info(
		"FindAccountsByUserIDService executed successfully",
		zap.String("journey", "findAccountsByuserID"),
	)

	return ad.accountRepository.FindAccountsByUserIDRepository(userID)
}

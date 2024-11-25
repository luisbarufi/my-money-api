package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"go.uber.org/zap"
)

func (td *transactionDomainService) FindTransactionsByUserIDService(
	userID uint64,
) ([]model.TransactionDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindTransactionsByUserIDService",
		zap.String("journey", "findTransactionsByUserID"),
	)

	logger.Info(
		"FindTransactionsByUserIDService executed successfully",
		zap.String("journey", "findTransactionsByUserID"),
	)

	return td.transactionRepository.FindTransactionsByUserIDRepository(userID)
}

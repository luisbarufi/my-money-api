package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"go.uber.org/zap"
)

func (td *transactionDomainService) ListTransactionsByUserIDService(
	userID uint64,
) ([]model.TransactionDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init ListTransactionsByUserIDService",
		zap.String("journey", "listTransactionsByUserID"),
	)

	logger.Info(
		"ListTransactionsByUserIDService executed successfully",
		zap.String("journey", "listTransactionsByUserID"),
	)

	return td.transactionRepository.ListTransactionsByUserIDRepository(userID)
}

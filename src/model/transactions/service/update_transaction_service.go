package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"go.uber.org/zap"
)

func (td *transactionDomainService) UpdateTransactionService(
	transactionID uint64, transactionDomain model.TransactionDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateTransactionService",
		zap.String("journey", "updateTransaction"),
	)

	err := td.transactionRepository.UpdateTransactionRepository(
		transactionID,
		transactionDomain,
	)

	if err != nil {
		logger.Error(
			"Error trying to call UpdateTransactionRepository",
			err, zap.String("journey", "updateTransaction"),
		)

		return err
	}

	logger.Info(
		"UpdateTransactionService executed successfully",
		zap.String("journey", "updateTransaction"),
	)

	return nil
}

package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (td *transactionDomainService) DeleteTransactionService(
	transactionID uint64,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteTransactionService",
		zap.String("journey", "deleteTransaction"),
	)

	err := td.transactionRepository.DeleteTransactionRepository(transactionID)

	if err != nil {
		logger.Error(
			"Error trying to call DeleteTransactionRepository",
			err,
			zap.String("journey", "deleteTransaction"),
		)

		return err
	}

	logger.Info(
		"DeleteTransactionService executed successfully",
		zap.String("transactionID", fmt.Sprintf("%d", transactionID)),
		zap.String("journey", "deleteTransaction"),
	)

	return nil
}

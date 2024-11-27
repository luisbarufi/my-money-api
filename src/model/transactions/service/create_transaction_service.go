package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"go.uber.org/zap"
)

func (td *transactionDomainService) CreateTransactionService(
	transactionDomain model.TransactionDomainInterface,
) (model.TransactionDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateTransactionService",
		zap.String("journey", "createTransaction"),
	)

	transactionDomainRepository, err := td.transactionRepository.CreateTransactionRepository(
		transactionDomain,
	)

	if err != nil {
		logger.Error(
			"Error trying to call CreateTransactionRepository",
			err,
			zap.String("journey", "createTransaction"),
		)

		return nil, err
	}

	logger.Info(
		"CreateTransactionService executed successfully",
		zap.String("userId", fmt.Sprintf("%d", transactionDomainRepository.GetID())),
		zap.String("journey", "createUser"),
	)

	return transactionDomainRepository, nil
}

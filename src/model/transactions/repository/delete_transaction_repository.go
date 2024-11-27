package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (tr *transactionRepository) DeleteTransactionRepository(
	transactionID uint64,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteTransactionRepository",
		zap.String("journey", "deleteTransaction"),
	)

	statement, err := tr.db.Prepare("DELETE FROM transactions WHERE id = $1")

	if err != nil {
		logger.Error(
			"Error preparing delete statement",
			err,
			zap.String("journey", "deleteTransaction"),
		)

		return nil
	}

	defer statement.Close()

	if _, err := statement.Exec(transactionID); err != nil {
		logger.Error(
			"Error executing delete statement",
			err,
			zap.String("journey", "deleteTransaction"),
		)

		return rest_err.NewInternalServerError("Error deleting transaction")
	}

	logger.Info(
		"DeleteTransactionRepository executed successfully",
		zap.String("journey", "deleteTransaction"),
	)

	return nil
}

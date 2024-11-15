package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ar *accountRepository) DeleteAccountRepository(
	accountId uint64,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteAccountRepository",
		zap.String("journey", "deleteAccount"),
	)

	statement, err := ar.db.Prepare("DELETE FROM accounts WHERE id = $1")

	if err != nil {
		logger.Error(
			"Error preparing delete statement",
			err,
			zap.String("journey", "deleteAccount"),
		)

		return nil
	}

	defer statement.Close()

	if _, err := statement.Exec(accountId); err != nil {
		logger.Error(
			"Error executing delete statement",
			err,
			zap.String("journey", "deleteAccount"),
		)

		return rest_err.NewInternalServerError("Error deleting account")
	}

	logger.Info(
		"DeleteAccountRepository executed successfully",
		zap.String("journey", "deleteAccount"),
	)

	return nil
}

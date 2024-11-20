package repository

import (
	"fmt"
	"time"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity/converter"
	"go.uber.org/zap"
)

func (ar *accountRepository) UpdateAccountRepository(
	accountId uint64, accountDomain model.AccountDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateAccountRepository",
		zap.String("journey", "updateAccount"),
	)

	statement, err := ar.db.Prepare(
		"UPDATE accounts SET account_name = $1, updated_at = $2 WHERE id = $3",
	)

	if err != nil {
		logger.Error(
			"Error preparing update statement",
			err,
			zap.String("journey", "updateAccount"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	value := converter.ConvertDomainToEntity(accountDomain)

	if _, err := statement.Exec(
		value.AccountName, time.Now().UTC(), accountId,
	); err != nil {
		logger.Error(
			"Error executing update statement",
			err,
			zap.String("journey", "updateAccount"),
		)

		return rest_err.NewInternalServerError("Error updating account")
	}

	logger.Info("UpdateAccountRepository successfully",
		zap.String("accountId", fmt.Sprintf("%d", accountId)),
		zap.String("journey", "updateAccount"),
	)

	return nil
}

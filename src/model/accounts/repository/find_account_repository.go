package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity/converter"
	"go.uber.org/zap"
)

func (ar *accountRepository) FindAccountsByUserIDRepository(userID uint64) (
	[]model.AccountDomainInterface, *rest_err.RestErr,
) {
	logger.Info(
		"Init FindAccountsByUserIDRepository",
		zap.String("journey", "findAccountsByUserID"),
	)

	row, err := ar.db.Query("SELECT * FROM accounts WHERE user_id = $1", userID)

	if err != nil {
		logger.Error(
			"Error executing find user query",
			err,
			zap.String("journey", "findUserByID"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer row.Close()

	var accounts []entity.AccountEntity

	for row.Next() {
		var account entity.AccountEntity
		if err := row.Scan(
			&account.ID,
			&account.UserID,
			&account.AccountName,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			logger.Error(
				"Error scanning insert account result",
				err,
				zap.String("journey", "createAccount"),
			)

			return nil, rest_err.NewInternalServerError(err.Error())
		}

		accounts = append(accounts, account)
	}

	return converter.ConvertEntitiesToDomains(accounts), nil
}

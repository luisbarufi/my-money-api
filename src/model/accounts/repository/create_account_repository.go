package repository

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/accounts/repository/entity/converter"
	"go.uber.org/zap"
)

func (ar *accountRepository) CreateAccountRepository(
	accountDomain model.AccountDomainInterface) (
	model.AccountDomainInterface, *rest_err.RestErr,
) {
	logger.Info(
		"Init CreateUserRepository",
		zap.String("journey", "createAccount"),
	)

	query := `INSERT INTO accounts (user_id, account_name, balance)
						VALUES ($1, $2, $3) 
						RETURNING id, user_id, account_name, balance, created_at, updated_at`

	value := converter.ConvertDomainToEntity(accountDomain)

	row, err := ar.db.Query(
		query,
		value.UserID,
		value.AccountName,
		value.Balance,
	)

	if err != nil {
		logger.Error(
			"Error executing insert account query",
			err,
			zap.String("journey", "createAccount"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer row.Close()

	var account entity.AccountEntity

	if row.Next() {
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
	}

	logger.Info(
		"CreateAccountRepository executed successfully",
		zap.String("userId", fmt.Sprintf("%d", account.ID)),
		zap.String("journey", "createAccount"),
	)

	return converter.ConvertEntityToDomain(account), nil
}

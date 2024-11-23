package repository

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity/converter"
	"go.uber.org/zap"
)

func (tr *transactionRepository) CreateTransactionRepository(
	transactionDomain model.TransactionDomainInterface,
) (model.TransactionDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateTransactionRepository",
		zap.String("journey", "createTransaction"),
	)

	query := `INSERT INTO transactions (user_id, account_id, category_id, amount, transaction_type, description)
						VALUES ($1, $2, $3, $4, $5, $6) 
						RETURNING id, user_id, account_id, category_id, amount, transaction_type, description, created_at, updated_at`

	value := converter.ConvertDomainToEntity(transactionDomain)

	row, err := tr.db.Query(
		query,
		value.UserID,
		value.AccountID,
		value.CategoryID,
		value.Amount,
		value.TransactionType,
		value.Description,
	)

	if err != nil {
		logger.Error(
			"Error executing insert transaction query",
			err,
			zap.String("journey", "createTransaction"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer row.Close()

	var transaction entity.TransactionEntity

	if row.Next() {
		if err := row.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.AccountID,
			&transaction.CategoryID,
			&transaction.Amount,
			&transaction.TransactionType,
			&transaction.Description,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		); err != nil {
			logger.Error(
				"Error scanning insert transaction result",
				err,
				zap.String("journey", "createTransaction"),
			)

			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	logger.Info(
		"CreateTransactionRepository executed successfully",
		zap.String("userId", fmt.Sprintf("%d", transaction.ID)),
		zap.String("journey", "createTransaction"),
	)

	return converter.ConvertEntityToDomain(transaction), nil
}

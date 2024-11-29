package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity/converter"
	"go.uber.org/zap"
)

func (tr *transactionRepository) ListTransactionsByUserIDRepository(
	userID uint64,
) ([]model.TransactionDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init ListTransactionsByUserIDRepository",
		zap.String("journey", "listTransactionByUserID"),
	)

	row, err := tr.db.Query("SELECT * FROM transactions WHERE user_id = $1", userID)

	if err != nil {
		logger.Error(
			"Error executing find transaction query",
			err,
			zap.String("journey", "listTransactionByUserID"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer row.Close()

	var transactions []entity.TransactionEntity

	for row.Next() {
		var transaction entity.TransactionEntity
		if err := row.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.AccountID,
			&transaction.CategoryID,
			&transaction.Amount,
			&transaction.Description,
			&transaction.TransactionType,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		); err != nil {
			logger.Error(
				"Error scanning insert category result",
				err,
				zap.String("journey", "listTransactionByUserID"),
			)

			return nil, rest_err.NewInternalServerError(err.Error())
		}

		transactions = append(transactions, transaction)
	}

	return converter.ConvertEntitiesToDomains(transactions), nil
}

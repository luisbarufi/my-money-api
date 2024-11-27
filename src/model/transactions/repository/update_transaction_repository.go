package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"github.com/luisbarufi/my-money-api/src/model/transactions/repository/entity/converter"
	"go.uber.org/zap"
)

func (tr *transactionRepository) UpdateTransactionRepository(
	transactionID uint64, transactionDomain model.TransactionDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateTransactionRepository",
		zap.String("journey", "updateTransaction"),
	)

	value := converter.ConvertDomainToEntity(transactionDomain)

	updates := make([]string, 0)

	args := make([]interface{}, 0)

	fieldsToUpdate := map[string]interface{}{
		"account_id":       value.AccountID,
		"category_id":      value.CategoryID,
		"amount":           value.Amount,
		"description":      value.Description,
		"transaction_type": value.TransactionType,
		"updated_at":       time.Now().UTC(),
	}

	for field, val := range fieldsToUpdate {
		if val != "" {
			updates = append(updates, fmt.Sprintf("%s = $%d", field, len(args)+1))

			args = append(args, val)
		}
	}

	args = append(args, transactionID)

	query := fmt.Sprintf(
		"UPDATE transactions SET %s WHERE id = $%d", strings.Join(updates, ", "),
		len(args),
	)

	if len(args) == 1 {
		logger.Info(
			"No parameters were provided to update",
			zap.String("journey", "updateTransaction"),
		)

		return nil
	}

	statement, err := tr.db.Prepare(query)

	if err != nil {
		logger.Error(
			"Error preparing update user query",
			err,
			zap.String("journey", "updateTransaction"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	if _, err := statement.Exec(args...); err != nil {
		logger.Error(
			"Error executing update transaction query",
			err,
			zap.String("journey", "updateTransaction"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"UpdateUserRepository successfully",
		zap.String("transactionID", fmt.Sprintf("%d", transactionID)),
		zap.String("journey", "updateTransaction"),
	)

	return nil
}

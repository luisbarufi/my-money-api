package repository

import (
	"fmt"
	"time"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity/converter"
	"go.uber.org/zap"
)

func (cr *categoryRepository) UpdateCategoryRepository(
	categoryId uint64, categoryDomain model.CategoryDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateCategoryRepository",
		zap.String("journey", "updateCategory"),
	)

	statement, err := cr.db.Prepare(
		"UPDATE categories SET category_name = $1, updated_at = $2 WHERE id = $3",
	)

	if err != nil {
		logger.Error(
			"Error preparing update statement",
			err,
			zap.String("journey", "updateCategory"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	value := converter.ConvertDomainToEntity(categoryDomain)

	if _, err := statement.Exec(
		value.CategoryName, time.Now().UTC(), categoryId,
	); err != nil {
		logger.Error(
			"Error executing update statement",
			err,
			zap.String("journey", "updateCategory"),
		)

		return rest_err.NewInternalServerError("Error updating category")
	}

	logger.Info(
		"UpdateAccountRepository successfully",
		zap.String("categoryId", fmt.Sprintf("%d", categoryId)),
		zap.String("journey", "updateCategory"),
	)

	return nil
}

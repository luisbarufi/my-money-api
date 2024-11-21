package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (cr *categoryRepository) DeleteCategoryRepository(
	categoryId uint64,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteCategoryRepository",
		zap.String("journey", "deleteCategory"),
	)

	statement, err := cr.db.Prepare("DELETE FROM categories WHERE id = $1")

	if err != nil {
		logger.Error(
			"Error preparing delete statement",
			err,
			zap.String("journey", "deleteCategory"),
		)

		return nil
	}

	defer statement.Close()

	if _, err := statement.Exec(categoryId); err != nil {
		logger.Error(
			"Error executing delete statement",
			err,
			zap.String("journey", "deleteCategory"),
		)

		return rest_err.NewInternalServerError("Error deleting category")
	}

	logger.Info(
		"DeleteCategoryRepository executed successfully",
		zap.String("journey", "deleteCategory"),
	)

	return nil
}

package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity/converter"
	"go.uber.org/zap"
)

func (cr *categoryRepository) ListCategoriesByUserIDRepository(
	userID uint64,
) ([]model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init ListCategoriesByUserIDRepository",
		zap.String("journey", "listCategoriesByUserID"),
	)

	row, err := cr.db.Query("SELECT * FROM categories WHERE user_id = $1", userID)

	if err != nil {
		logger.Error(
			"Error executing find category query",
			err,
			zap.String("journey", "listCategoriesByUserID"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer row.Close()

	var categories []entity.CategoryEntity

	for row.Next() {
		var category entity.CategoryEntity
		if err := row.Scan(
			&category.ID,
			&category.UserID,
			&category.CategoryName,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			logger.Error(
				"Error scanning insert category result",
				err,
				zap.String("journey", "listCategoriesByUserID"),
			)

			return nil, rest_err.NewInternalServerError(err.Error())
		}

		categories = append(categories, category)
	}

	return converter.ConvertEntitiesToDomains(categories), nil
}

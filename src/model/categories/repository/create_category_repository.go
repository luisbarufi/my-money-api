package repository

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/categories/repository/entity/converter"
	"go.uber.org/zap"
)

func (cr *categoryRepository) CreateCategoryRepository(
	categoryDomain model.CategoryDomainInterface,
) (model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateCategoryRepository",
		zap.String("journey", "createCategory"),
	)

	query := `INSERT INTO categories (user_id, category_name)
						VALUES ($1, $2) 
						RETURNING id, user_id, category_name, created_at, updated_at`

	value := converter.ConvertDomainToEntity(categoryDomain)

	row, err := cr.db.Query(
		query,
		value.UserID,
		value.CategoryName,
	)

	if err != nil {
		logger.Error(
			"Error executing insert category query",
			err,
			zap.String("journey", "createCategory"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer row.Close()

	var category entity.CategoryEntity

	if row.Next() {
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
				zap.String("journey", "createCategory"),
			)

			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	logger.Info(
		"CreateAccountRepository executed successfully",
		zap.String("categoryID", fmt.Sprintf("%d", category.ID)),
		zap.String("journey", "createAccount"),
	)

	return converter.ConvertEntityToDomain(category), nil
}

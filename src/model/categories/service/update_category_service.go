package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"go.uber.org/zap"
)

func (cd *categoryDomainService) UpdateCategoryService(
	categoryId uint64, categoryDomain model.CategoryDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateCategoryService",
		zap.String("journey", "updateCategory"),
	)

	err := cd.categoryRepository.UpdateCategoryRepository(
		categoryId,
		categoryDomain,
	)

	if err != nil {
		logger.Error(
			"Error trying to call UpdateCategoryRepository",
			err, zap.String("journey", "updateCategory"),
		)

		return err
	}

	logger.Info(
		"UpdateCategoryService executed successfully",
		zap.String("journey", "updateCategory"),
	)

	return nil
}

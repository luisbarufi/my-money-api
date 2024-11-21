package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (cd *categoryDomainService) DeleteCategoryService(
	categoryId uint64,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteCategoryService",
		zap.String("journey", "deleteCategory"),
	)

	err := cd.categoryRepository.DeleteCategoryRepository(categoryId)

	if err != nil {
		logger.Error(
			"Error trying to call DeleteCategoryRepository",
			err,
			zap.String("journey", "deleteCategory"),
		)

		return err
	}

	logger.Info(
		"DeleteCategoryService executed successfully",
		zap.String("categoryId", fmt.Sprintf("%d", categoryId)),
		zap.String("journey", "deleteCategory"),
	)

	return nil
}

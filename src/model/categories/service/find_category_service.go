package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"go.uber.org/zap"
)

func (cd *categoryDomainService) FindCategoriesByUserIDService(
	userID uint64,
) ([]model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindCategoriesByUserIDService",
		zap.String("journey", "findCategoriesByUserID"),
	)

	logger.Info(
		"FindCategoriesByUserIDService executed successfully",
		zap.String("journey", "findCategoriesByUserID"),
	)

	return cd.categoryRepository.FindCategoriesByUserIDRepository(userID)
}

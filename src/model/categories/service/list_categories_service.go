package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"go.uber.org/zap"
)

func (cd *categoryDomainService) ListCategoriesByUserIDService(
	userID uint64,
) ([]model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init ListCategoriesByUserIDService",
		zap.String("journey", "listCategoriesByUserID"),
	)

	logger.Info(
		"ListCategoriesByUserIDService executed successfully",
		zap.String("journey", "listCategoriesByUserID"),
	)

	return cd.categoryRepository.ListCategoriesByUserIDRepository(userID)
}

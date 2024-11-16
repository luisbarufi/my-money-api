package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"go.uber.org/zap"
)

func (cd *categoryDomainService) CreateCategoryService(
	categoryDomain model.CategoryDomainInterface,
) (model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateCategoryService",
		zap.String("journey", "createCategory"),
	)

	categoryDomainRepository, err := cd.categoryRepository.CreateCategoryRepository(categoryDomain)

	if err != nil {
		logger.Error("Error trying to call CreateCategoryRepository",
			err,
			zap.String("journey", "createCategory"),
		)

		return nil, err
	}

	logger.Info(
		"CreateCategoryService executed successfully",
		zap.String("userId", fmt.Sprintf("%d", categoryDomainRepository.GetID())),
		zap.String("journey", "createUser"),
	)

	return categoryDomainRepository, nil
}

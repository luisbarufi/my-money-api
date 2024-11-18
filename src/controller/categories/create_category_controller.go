package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/categories/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	view "github.com/luisbarufi/my-money-api/src/view/categories"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) CreateCategoryController(c *gin.Context) {
	logger.Info("Init CreateCategoryController",
		zap.String("journey", "createCategory"),
	)

	var categoryRequest request.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		logger.Error(
			"Error trying to validate category info",
			err,
			zap.String("journey", "createCategory"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	secretKey := env.GetEnv("JWT_SECRET_KEY")

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	claims, err := utils.ParseToken(token, secretKey)
	if err != nil {
		logger.Error(
			"Error when trying to parse to token",
			err,
			zap.String("journey", "createCategory"),
		)

		c.JSON(err.Code, err)

		return
	}

	userID, err := utils.ExtractUserID(claims)

	if err != nil {
		logger.Error(
			"Error while trying to extract token ID",
			err,
			zap.String("journey", "createCategory"),
		)

		c.JSON(err.Code, err)

		return
	}

	domain := model.NewCategoryDomain(
		userID,
		categoryRequest.CategoryName,
	)

	domainResult, err := cc.categoryService.CreateCategoryService(domain)

	if err != nil {
		logger.Error(
			"Error calling CreateCategoryService",
			err,
			zap.String("journey", "createCategory"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"CreateCategoryController executed successfully",
		zap.String("journey", "createCategory"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
	view "github.com/luisbarufi/my-money-api/src/view/categories"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) ListCategoriesByUserIDController(
	c *gin.Context,
) {
	logger.Info(
		"Init ListCategoriesByUserIDController",
		zap.String("journey", "listCategoriesByUserID"),
	)

	secretKey := env.GetEnv("JWT_SECRET_KEY")

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	claims, err := utils.ParseToken(token, secretKey)
	if err != nil {
		logger.Error(
			"Error when trying to parse to token",
			err,
			zap.String("journey", "listCategoriesByUserID"),
		)

		c.JSON(err.Code, err)

		return
	}

	userID, err := utils.ExtractUserID(claims)

	if err != nil {
		logger.Error(
			"Error while trying to extract token ID",
			err,
			zap.String("journey", "listCategoriesByUserID"),
		)

		c.JSON(err.Code, err)

		return
	}

	categoriesDomain, restErr := cc.categoryService.ListCategoriesByUserIDService(userID)

	if restErr != nil {
		logger.Error("Error calling ListCategoriesByUserIDService",
			restErr,
			zap.String("journey", "listCategoriesByUserID"),
		)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"ListCategoriesByUserIDController executed successfully",
		zap.String("journey", "findAccountsByuserID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainsToResponses(categoriesDomain))
}

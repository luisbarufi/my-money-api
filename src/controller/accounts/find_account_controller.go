package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
	view "github.com/luisbarufi/my-money-api/src/view/accounts"
	"go.uber.org/zap"
)

func (ac *accountControllerInterface) FindAccountsByUserIDController(c *gin.Context) {
	logger.Info(
		"Init FindAccountsByUserIDController",
		zap.String("journey", "findAccountsByUserID"),
	)

	secretKey := env.GetEnv("JWT_SECRET_KEY")

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	claims, err := utils.ParseToken(token, secretKey)
	if err != nil {
		logger.Error(
			"Error when trying to parse to token",
			err,
			zap.String("journey", "findAccountsByUserID"),
		)

		c.JSON(err.Code, err)

		return
	}

	userID, err := utils.ExtractUserID(claims)

	if err != nil {
		logger.Error(
			"Error while trying to extract token ID",
			err,
			zap.String("journey", "findAccountsByUserID"),
		)

		c.JSON(err.Code, err)

		return
	}

	accountDomain, restErr := ac.accountService.FindAccountsByUserIDService(userID)

	if restErr != nil {
		logger.Error("Error calling FindAccountsByUserIDService",
			restErr,
			zap.String("journey", "findAccountsByUserID"),
		)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"FindAccountsByUserIDController executed successfully",
		zap.String("journey", "findAccountsByUserID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainsToResponses(accountDomain))
}

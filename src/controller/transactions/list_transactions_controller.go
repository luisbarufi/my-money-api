package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
	view "github.com/luisbarufi/my-money-api/src/view/transactions"
	"go.uber.org/zap"
)

func (tc *transactionControllerInterface) ListTrasactionsByUserIDController(
	c *gin.Context,
) {
	logger.Info(
		"Init ListTrasactionsByUserIDController",
		zap.String("journey", "listTransactionByUserID"),
	)

	secretKey := env.GetEnv("JWT_SECRET_KEY")

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	claims, err := utils.ParseToken(token, secretKey)
	if err != nil {
		logger.Error(
			"Error when trying to parse to token",
			err,
			zap.String("journey", "listTransactionByUserID"),
		)

		c.JSON(err.Code, err)

		return
	}

	userID, err := utils.ExtractUserID(claims)

	if err != nil {
		logger.Error(
			"Error while trying to extract token ID",
			err,
			zap.String("journey", "listTransactionByUserID"),
		)

		c.JSON(err.Code, err)

		return
	}

	transactionDomain, restErr := tc.transactionService.ListTransactionsByUserIDService(userID)

	if restErr != nil {
		logger.Error("Error calling ListTransactionsByUserIDService",
			restErr,
			zap.String("journey", "listTransactionByUserID"),
		)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"ListTrasactionsByUserIDController executed successfully",
		zap.String("journey", "listTransactionByUserID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainsToResponses(transactionDomain))
}

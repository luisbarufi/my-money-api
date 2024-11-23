package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/transactions/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	view "github.com/luisbarufi/my-money-api/src/view/transactions"
	"go.uber.org/zap"
)

func (tc *transactionControllerInterface) CreateTransactionController(
	c *gin.Context,
) {
	logger.Info(
		"Init CreateTransactionController",
		zap.String("journey", "createTransaction"),
	)

	var transactionRequest request.TransactionRequest

	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		logger.Error(
			"Error trying to validate transaction info",
			err,
			zap.String("journey", "createTransaction"),
		)

		restErr := validation.ValidateError(err)

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
			zap.String("journey", "createTransaction"),
		)

		c.JSON(err.Code, err)

		return
	}

	userID, err := utils.ExtractUserID(claims)

	if err != nil {
		logger.Error(
			"Error while trying to extract token ID",
			err,
			zap.String("journey", "createTransaction"),
		)

		c.JSON(err.Code, err)

		return
	}

	domain := model.NewTrasactionDomain(
		userID,
		transactionRequest.AccountID,
		transactionRequest.CategoryID,
		transactionRequest.Amount,
		transactionRequest.Description,
		transactionRequest.TransactionType,
	)

	domainResult, err := tc.transactionService.CreateTransactionService(domain)

	if err != nil {
		logger.Error(
			"Error calling CreateTransactionService",
			err,
			zap.String("journey", "createTransaction"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"CreateTransactionController executed successfully",
		zap.String("journey", "createTransaction"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

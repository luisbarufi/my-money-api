package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/transactions/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
	"go.uber.org/zap"
)

func (tc *transactionControllerInterface) UpdateTransactionController(
	c *gin.Context,
) {
	logger.Info(
		"Init UpdateTransactionController",
		zap.String("journey", "updateTransaction"),
	)

	transactionId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(
			"Error trying to validate transaction id, must be integer",
			err,
			zap.String("journey", "updateTransaction"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	var transactionRequest request.TransactionRequest

	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		logger.Error(
			"Error trying to validate transaction info",
			err,
			zap.String("journey", "updateTransaction"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUpdateTransactionDomain(
		transactionRequest.AccountID,
		transactionRequest.CategoryID,
		transactionRequest.Amount,
		transactionRequest.Description,
		transactionRequest.TransactionType,
	)

	if err := tc.transactionService.UpdateTransactionService(transactionId, domain); err != nil {
		logger.Error("Error calling UpdateTransactionService",
			err,
			zap.String("journey", "updateTransaction"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"UpdateTransactionController executed successfully",
		zap.String("transactionId", fmt.Sprintf("%d", transactionId)),
		zap.String("journey", "updateTransaction"),
	)

	c.Status(http.StatusOK)
}

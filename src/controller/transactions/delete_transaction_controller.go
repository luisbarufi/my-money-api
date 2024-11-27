package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"go.uber.org/zap"
)

func (tc *transactionControllerInterface) DeleteTransactionController(
	c *gin.Context,
) {
	logger.Info(
		"Init DeleteTransactionController",
		zap.String("journey", "deleteTransaction"),
	)

	transactionId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(
			"Error trying to validate transaction id, must be integer",
			err,
			zap.String("journey", "deleteTransaction"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	if err := tc.transactionService.DeleteTransactionService(transactionId); err != nil {
		logger.Error(
			"Error calling DeleteTransactionService",
			err,
			zap.String("journey", "deleteTransaction"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"DeleteTransactionController executed successfully",
		zap.String("transactionId", fmt.Sprintf("%d", transactionId)),
		zap.String("journey", "deleteTransaction"),
	)

	c.Status(http.StatusOK)
}

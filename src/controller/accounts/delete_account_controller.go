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

func (ac *accountControllerInterface) DeleteAccountController(c *gin.Context) {
	logger.Info(
		"Init DeleteAccountController",
		zap.String("journey", "deleteAccount"),
	)

	accountId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(
			"Error trying to validate account id, must be integer",
			err,
			zap.String("journey", "deleteAccount"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	if err := ac.accountService.DeleteAccountService(accountId); err != nil {
		logger.Error(
			"Error calling DeleteAccountService",
			err,
			zap.String("journey", "deleteAccount"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info("DeleteAccountController executed successfully",
		zap.String("accountId", fmt.Sprintf("%d", accountId)),
		zap.String("journey", "deleteAccount"),
	)

	c.Status(http.StatusOK)
}

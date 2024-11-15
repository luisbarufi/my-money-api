package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/accounts/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	"go.uber.org/zap"
)

func (ac *accountControllerInterface) UpdateAccountController(c *gin.Context) {
	logger.Info(
		"Init UpdateAccountController",
		zap.String("journey", "updateAccount"),
	)

	accountId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(
			"Error trying to validate account id, must be integer",
			err,
			zap.String("journey", "updateAccount"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	var accountRequest request.UpdateAccountRequest

	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "updateAccount"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewAccountUpdateDomain(accountRequest.AccountName)

	if err := ac.accountService.UpdateAccountService(accountId, domain); err != nil {
		logger.Error("Error calling UpdateUserService",
			err,
			zap.String("journey", "updateAccount"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"UpdateAccountController executed successfully",
		zap.String("userId", fmt.Sprintf("%d", accountId)),
		zap.String("journey", "updateAccount"),
	)

	c.Status(http.StatusOK)
}

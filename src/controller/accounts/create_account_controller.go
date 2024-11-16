package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/accounts/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
	view "github.com/luisbarufi/my-money-api/src/view/accounts"
	"go.uber.org/zap"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ac *accountControllerInterface) CreateAccountController(c *gin.Context) {
	logger.Info("Init CreateAccountController",
		zap.String("journey", "createAccount"),
	)

	var accountRequest request.AccountRequest

	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		logger.Error(
			"Error trying to validate account info",
			err,
			zap.String("journey", "createAccount"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	secretKey := env.GetEnv(JWT_SECRET_KEY)

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	claims, err := utils.ParseToken(token, secretKey)
	if err != nil {
		logger.Error(
			"Error when trying to parse to token",
			err,
			zap.String("journey", "createAccount"),
		)

		c.JSON(err.Code, err)

		return
	}

	userID, err := utils.ExtractUserID(claims)

	if err != nil {
		logger.Error(
			"Error while trying to extract token ID",
			err,
			zap.String("journey", "createAccount"),
		)

		c.JSON(err.Code, err)

		return
	}

	domain := model.NewAccountDomain(
		userID,
		accountRequest.AccountName,
		accountRequest.Balance,
	)

	domainResult, err := ac.accountService.CreateAccountService(domain)

	if err != nil {
		logger.Error(
			"Error calling CreateAccountService",
			err,
			zap.String("journey", "createAccount"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"CreateAccountController executed successfully",
		zap.String("journey", "createAccount"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/users/model/request"
	"github.com/luisbarufi/my-money-api/src/controller/users/model/response"
	model "github.com/luisbarufi/my-money-api/src/model/users"
	view "github.com/luisbarufi/my-money-api/src/view/users"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) ForgotPasswordController(c *gin.Context) {
	logger.Info(
		"Init ForgotPasswordController",
		zap.String("journey", "forgotPassword"),
	)

	var userForgotPassword request.UserForgotPassword

	if err := c.ShouldBindJSON(&userForgotPassword); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "forgotPassword"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserForgotPasswordDomain(
		userForgotPassword.Email,
	)

	domainResult, resetToken, err := uc.service.ForgotPasswordService(domain)

	if err != nil {
		logger.Error(
			"Error calling ForgotPasswordService",
			err,
			zap.String("journey", "forgotPassword"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"ForgotPasswordController executed successfully",
		zap.String("userId", fmt.Sprintf("%d", domainResult.GetID())),
		zap.String("journey", "forgotPassword"),
	)

	response := response.PasswordForgotResponse{
		User:       view.ConvertDomainToResponse(domainResult),
		ResetToken: resetToken,
	}

	c.JSON(http.StatusOK, response)
}

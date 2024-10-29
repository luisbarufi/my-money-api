package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/model/request"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/view"
	"go.uber.org/zap"
)

type PasswordResetResponse struct {
	User       interface{} `json:"user"`
	ResetToken string      `json:"resetToken"`
}

func (uc *userControllerInterface) PasswordResetController(c *gin.Context) {

	var userResetPassword request.UserResetPassword

	if err := c.ShouldBindJSON(&userResetPassword); err != nil {
		logger.Error(
			"Error trying to validate passwordReset info",
			err,
			zap.String("journey", "passwordReset"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserResetPasswordDomain(
		userResetPassword.Email,
		userResetPassword.NewPassword,
	)

	domainResult, resetToken, err := uc.service.ResetPasswordService(domain)
	if err != nil {
		logger.Error(
			"Error calling passwordReset service",
			err,
			zap.String("journey", "passwordReset"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("PasswordReset controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", domainResult.GetID())),
		zap.String("journey", "passwordReset"),
	)

	response := PasswordResetResponse{
		User:       view.ConvertDomainToResponse(domainResult),
		ResetToken: resetToken,
	}

	c.JSON(http.StatusOK, response)
}

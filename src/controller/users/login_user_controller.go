package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/users/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/users"
	view "github.com/luisbarufi/my-money-api/src/view/users"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUserController(c *gin.Context) {
	logger.Info("Init LoginUserController", zap.String("journey", "loginUser"))

	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "loginUser"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)

	if err != nil {
		logger.Error(
			"Error calling LoginUserService",
			err,
			zap.String("journey", "loginUser"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"LoginUserController executed successfully",
		zap.String("userId", fmt.Sprintf("%d", domainResult.GetID())),
		zap.String("journey", "loginUser"),
	)

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

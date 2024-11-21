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

func (uc *userControllerInterface) CreateUserController(c *gin.Context) {
	logger.Info("Init CreateUserController", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "createUser"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Nick,
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := uc.service.CreateUserService(domain)

	if err != nil {
		logger.Error(
			"Error calling CreateUserService",
			err,
			zap.String("journey", "createUser"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"CreateUserController executed successfully",
		zap.String("userId", fmt.Sprintf("%d", domainResult.GetID())),
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

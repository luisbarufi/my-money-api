package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUserController(c *gin.Context) {
	logger.Info("Init createUser controller", zap.String("journey", "createUser"))

	userRequest, restErr := validation.ValidateUserInput(c)
	if restErr != nil {
		logger.Error(
			"Error calling ValidateUserInput",
			restErr,
			zap.String("journey", "createUser"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Nick,
		userRequest.Email,
		userRequest.Password,
	)
	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error(
			"Error calling createUser service",
			err,
			zap.String("journey", "createUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", domainResult.GetID())),
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

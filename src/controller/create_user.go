package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/model/request"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller", zap.String("Journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validation user info", err, zap.String("Journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User Successfully Created", zap.String("Journey", "createUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

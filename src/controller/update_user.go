package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/model/request"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))

	var userRequest request.UserUpdateRequest

	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil ||
		strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validation user info",
			err,
			zap.String("journey", "updateUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Nick,
	)
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser services",
			err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"updateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	c.Status(http.StatusOK)
}

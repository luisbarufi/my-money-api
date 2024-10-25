package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUserController(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))

	userId, restErr := validation.ValidateUserID(c)
	if restErr != nil {
		logger.Error(
			"Error calling ValidateUserID",
			restErr,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error(
			"Error calling deleteUser service",
			err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "deleteUser"),
	)
	c.Status(http.StatusOK)
}

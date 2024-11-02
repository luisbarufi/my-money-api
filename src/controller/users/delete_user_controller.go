package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUserController(c *gin.Context) {
	logger.Info("Init DeleteUserController", zap.String("journey", "deleteUser"))

	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error("Error trying to validate user id, must be integer",
			err,
			zap.String("journey", "deleteUser"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	if err := uc.service.DeleteUserService(userId); err != nil {
		logger.Error(
			"Error calling DeleteUserService",
			err,
			zap.String("journey", "deleteUser"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info("DeleteUserController executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}

package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))

	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Error trying to validate user id, must be integer",
			err,
			zap.String("journey", "deleteUser"),
		)
		errorMessage := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	uc.service.DeleteUser(userId)

	logger.Info(
		"deleteUser controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "deleteUser"),
	)
	c.Status(http.StatusOK)
}

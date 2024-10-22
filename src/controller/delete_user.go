package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))

	userId, restErr := validation.ValidateUserID(c)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	if err := uc.callDeleteUserService(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "deleteUser"),
	)
	c.Status(http.StatusOK)
}

func (uc *userControllerInterface) callDeleteUserService(
	userId uint64,
) *rest_err.RestErr {
	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error calling deleteUser service",
			err, zap.String("journey", "deleteUser"),
		)
		return err
	}
	return nil
}

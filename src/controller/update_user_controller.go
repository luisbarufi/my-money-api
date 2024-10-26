package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUserController(c *gin.Context) {
	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))

	userId, restErr := validation.ValidateUserID(c)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	userRequest, restErr := validation.ValidateUserUpdateRequest(c)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Nick)
	err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error("Error calling UpdateUser service",
			err, zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "updateUser"),
	)
	c.Status(http.StatusOK)
}

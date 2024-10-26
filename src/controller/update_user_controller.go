package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/model/request"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUserController(c *gin.Context) {
	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))

	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Error trying to validate user id, must be integer",
			err,
			zap.String("journey", "findUserByID"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	var userRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Nick)
	if err := uc.service.UpdateUserService(userId, domain); err != nil {
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

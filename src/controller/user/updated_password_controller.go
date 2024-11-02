package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/user/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/user"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdatePasswordController(c *gin.Context) {
	logger.Info("Init UpdatePasswordController",
		zap.String("journey", "updatepassword"),
	)

	var userUpdatePassword request.UserUpdatePassword
	if err := c.ShouldBindJSON(&userUpdatePassword); err != nil {
		logger.Error(
			"Error trying to validate userUpdatePassword info",
			err,
			zap.String("journey", "updatepassword"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	token := userUpdatePassword.Token
	domain := model.NewUserUpdatePasswordDomain(userUpdatePassword.Password)

	if err := uc.service.UpdatePasswordService(token, domain); err != nil {
		logger.Error("Error calling UpdatePasswordService",
			err, zap.String("journey", "updatepassword"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdatePasswordController executed successfully",
		zap.String("journey", "updatepassword"),
	)
	c.Status(http.StatusOK)
}

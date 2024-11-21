package controller

import (
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	view "github.com/luisbarufi/my-money-api/src/view/users"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByIDController(c *gin.Context) {
	logger.Info(
		"Init FindUserByIDController",
		zap.String("journey", "findUserByID"),
	)

	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error("Error trying to validate user id, must be integer",
			err,
			zap.String("journey", "findUserByID"),
		)

		restErr := rest_err.NewBadRequestError("Invalid user id")

		c.JSON(restErr.Code, restErr)

		return
	}

	userDomain, restErr := uc.service.FindUserByIDService(userId)

	if restErr != nil {
		logger.Error("Error calling FindUserByIDService",
			restErr,
			zap.String("journey", "findUserByID"),
		)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"FindUserByIDController executed successfully",
		zap.String("journey", "findUserByID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmailController(c *gin.Context) {
	logger.Info("Init FindUserByEmailController",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("email")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error(
			"Error trying to validate email",
			err,
			zap.String("journey", "findUserByEmail"),
		)

		restErr := validation.ValidateError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)

	if err != nil {
		logger.Error(
			"Error calling FindUserByEmailService",
			err,
			zap.String("journey", "findUserByEmail"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"FindUserByEmailController executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

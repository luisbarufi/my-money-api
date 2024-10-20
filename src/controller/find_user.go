package controller

import (
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info(
		"Init findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("Error trying to validade userId",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError("ID is not a valid")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, _ := uc.service.FindUserByIDServices(userID)
	if userDomain.GetID() == 0 {
		logger.Error(
			"Error trying to call findUserByID services",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewNotFoundError("User not found")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	logger.Info(
		"findUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info(
		"Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)
	userEmail := c.Param("email")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error(
			"Error trying to validade email",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError("Email is not a valid")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error(
			"Error trying to call findUserByEmail services",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}

	if userDomain.GetID() == 0 {
		logger.Info(
			"User not found",
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewNotFoundError("User not found")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	logger.Info(
		"findUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

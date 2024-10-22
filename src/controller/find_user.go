package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userID, restErr := validation.ValidateUserID(c)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain, err := uc.callFindUserByIDService(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) callFindUserByIDService(userID uint64) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	userDomain, err := uc.service.FindUserByIDServices(userID)
	if err != nil {
		logger.Error("Error calling FindUserByID service",
			err, zap.String("journey", "findUserByID"),
		)
		return nil, err
	}
	if userDomain.GetID() == 0 {
		return nil, rest_err.NewNotFoundError("User not found")
	}
	return userDomain, nil
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail, restErr := validation.ValidateUserEmail(c)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain, err := uc.callFindUserByEmailService(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) callFindUserByEmailService(
	userEmail string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error calling FindUserByEmail service",
			err, zap.String("journey", "findUserByEmail"),
		)
		return nil, err
	}
	if userDomain.GetID() == 0 {
		return nil, rest_err.NewNotFoundError("User not found")
	}
	return userDomain, nil
}

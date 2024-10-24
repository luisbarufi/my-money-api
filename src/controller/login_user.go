package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/model/request"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller", zap.String("journey", "loginUser"))

	userLogin, restErr := validation.ValidateLoginUserInput(c)
	if restErr != nil {
		logger.Error(
			"Error calling ValidateLoginUserInput",
			restErr,
			zap.String("journey", "loginUser"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	domainResult, err := uc.callLoginUserService(userLogin)
	if err != nil {
		logger.Error(
			"Error calling callLoginUserService",
			err,
			zap.String("journey", "loginUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("loginUser controller executed successfully",
		zap.String("userId", fmt.Sprintf("%d", domainResult.GetID())),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

func (uc *userControllerInterface) callLoginUserService(
	userLogin *request.UserLogin,
) (model.UserDomainInterface, *rest_err.RestErr) {
	domain := model.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error(
			"Error calling loginUser service",
			err,
			zap.String("journey", "loginUser"),
		)
		return nil, err
	}

	return domainResult, nil
}

package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) ResetPasswordService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init reset password services", zap.String("journey", "resetPasswordService"))

	user, err := ud.FindUserByEmailService(userDomain.GetEmail())
	if err != nil {
		logger.Error(
			"Error calling findUserByEmailAndPasswordServices",
			err,
			zap.String("journey", "resetPasswordService"),
		)
		return nil, "", err
	}

	resetToken, err := user.GenerateResetToken()
	if err != nil {
		logger.Error(
			"Error generating jwt token",
			err,
			zap.String("journey", "resetPasswordService"),
		)
		return nil, "", err
	}

	logger.Info(
		"ResetPasswordService services executed successfully",
		zap.String("userId", fmt.Sprintf("%d", user.GetID())),
		zap.String("journey", "resetPasswordService"),
	)
	return user, resetToken, nil
}

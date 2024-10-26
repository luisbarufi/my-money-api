package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init loginUser services", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()
	user, err := ud.findUserByEmailAndPasswordService(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		logger.Error(
			"Error calling findUserByEmailAndPasswordServices",
			err,
			zap.String("journey", "loginUser"),
		)
		return nil, err
	}

	logger.Info(
		"LoginUser services executed successfully",
		zap.String("userId", fmt.Sprintf("%d", user.GetID())),
		zap.String("journey", "loginUser"),
	)
	return user, nil
}

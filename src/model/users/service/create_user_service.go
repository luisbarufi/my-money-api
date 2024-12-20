package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/users"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUserService", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())

	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already in use")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUserRepository(userDomain)

	if err != nil {
		logger.Error(
			"Error trying to call repository",
			err,
			zap.String("journey", "createUser"),
		)

		return nil, err
	}

	logger.Info(
		"CreateUserService executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userDomainRepository.GetID())),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}

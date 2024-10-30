package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(
	userId uint64, userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("UpdateUserService", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUserRepository(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call UpdateUserRepository",
			err, zap.String("journey", "updateUser"),
		)
		return err
	}

	logger.Info(
		"UpdateUserService executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "updateUser"),
	)
	return nil
}

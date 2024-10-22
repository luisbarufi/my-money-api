package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId uint64) *rest_err.RestErr {
	logger.Info("Init deleteUser services", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		logger.Error(
			"Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"),
		)
		return err
	}

	logger.Info(
		"deleteUser service executed successfully",
		zap.String("userId", fmt.Sprintf("%d", userId)),
		zap.String("journey", "deleteUser"),
	)
	return nil
}

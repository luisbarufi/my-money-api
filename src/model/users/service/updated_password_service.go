package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/users"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdatePasswordService(
	token string, userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info(
		"Init UpdatePasswordService",
		zap.String("journey", "updatePassword"),
	)

	userDomain.EncryptPassword()

	err := ud.userRepository.UpdatePasswordRepository(token, userDomain)
	if err != nil {
		logger.Error(
			"Error trying to call UpdatePasswordRepository",
			err, zap.String("journey", "updatePassword"),
		)

		return err
	}

	logger.Info(
		"UpdatePasswordService executed successfully",
		zap.String("journey", "updatePassword"),
	)

	return nil
}

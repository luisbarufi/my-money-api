package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/users"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDService(id uint64) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init FindUserByIDService", zap.String("journey", "findUserByID"))

	logger.Info(
		"FindUserByIDService executed successfully",
		zap.String("journey", "findUserByID"),
	)

	return ud.userRepository.FindUserByIDRepository(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info(
		"Init FindUserByEmailService",
		zap.String("journey", "findUserByEmail"),
	)

	logger.Info(
		"FindUserByEmailService executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	return ud.userRepository.FindUserByEmailRepository(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(
	email, password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init findUserByEmailAndPasswordService",
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	logger.Info(
		"findUserByEmailAndPasswordService executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	return ud.userRepository.FindUserByEmailAndPasswordRepository(email, password)
}

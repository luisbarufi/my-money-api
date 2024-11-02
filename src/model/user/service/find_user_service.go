package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDService(id uint64) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByID", zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByIDRepository(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByEmail", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailRepository(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(
	email, password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailAndPasswordRepository(email, password)
}

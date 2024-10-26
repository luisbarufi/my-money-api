package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id uint64) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByID", zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByIDRepository(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByEmail", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailRepository(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email, password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailAndPasswordRepository(email, password)
}

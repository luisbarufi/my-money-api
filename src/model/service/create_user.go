package service

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomainService model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser Model", zap.String("Journey", "createUser"))

	userDomainService.EncryptPassword()

	fmt.Println(userDomainService.GetPassword())

	return nil
}

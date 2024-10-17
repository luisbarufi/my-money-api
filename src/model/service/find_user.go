package service

import (
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
)

func (ud *userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr,
) {
	return nil, nil
}

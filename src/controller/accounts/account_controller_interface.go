package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/model/accounts/service"
)

func NewAccountControllerInterface(
	accountServiceInterface service.AccountDomainService,
) AccountControllerInterface {
	return &accountControllerInterface{
		accountService: accountServiceInterface,
	}
}

type AccountControllerInterface interface {
	CreateAccountController(c *gin.Context)
}

type accountControllerInterface struct {
	accountService service.AccountDomainService
}

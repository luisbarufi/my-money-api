package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUserController(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUserController(c *gin.Context)

	FindUserByIDController(c *gin.Context)
	FindUserByEmailController(c *gin.Context)

	LoginUserController(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
)

func NewUserDomain(
	name, email, password string,
) UserDomainInterface {
	return &UserDomain{
		name, email, password,
	}
}

type UserDomain struct {
	Name     string
	Email    string
	Password string
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUserByID(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

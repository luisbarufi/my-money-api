package model

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	SECRET_KEY     = "SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := env.GetEnv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"name":  ud.name,
		"nick":  ud.nick,
		"email": ud.email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()),
		)
	}

	return tokenString, nil
}

func (ud *userDomain) GenerateResetToken() (string, *rest_err.RestErr) {

	secretKey := env.GetEnv(SECRET_KEY)

	claims := jwt.MapClaims{
		"id":  ud.id,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()),
		)
	}

	return signedToken, nil
}

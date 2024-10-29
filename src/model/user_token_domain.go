package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
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
		"user_id": ud.id,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
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

func ParseAndValidateResetToken(tokenString, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido ou expirado")
	}

	return token, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	secret := env.GetEnv(JWT_SECRET_KEY)
	tokenValue := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

	token, err := jwt.Parse(strings.TrimPrefix(tokenValue, "Bearer "),
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, rest_err.NewBadRequestError("Invalid token")
		},
	)

	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := userDomain{
		id:    uint64(claims["id"].(float64)),
		name:  claims["name"].(string),
		nick:  claims["nick"].(string),
		email: claims["email"].(string),
	}

	logger.Info(fmt.Sprintf("User authenticated: %v", userDomain))
}

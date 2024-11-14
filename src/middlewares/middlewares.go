package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

type middlewareUserDomain struct {
	id    uint64
	name  string
	nick  string
	email string
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

	userDomain := middlewareUserDomain{
		id:    uint64(claims["id"].(float64)),
		name:  claims["name"].(string),
		nick:  claims["nick"].(string),
		email: claims["email"].(string),
	}

	logger.Info(fmt.Sprintf("User authenticated: %v", userDomain))
}

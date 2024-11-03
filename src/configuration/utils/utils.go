package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
)

func ParseToken(tokenString, secretKey string) (jwt.MapClaims, *rest_err.RestErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		logger.Error("Error trying to parse token", err)
		errRest := rest_err.NewBadRequestError("Invalid token")

		return nil, errRest
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	logger.Error("Error trying to validate token", err)

	return nil, rest_err.NewBadRequestError("Invalid token")
}

func ExtractUserID(claims jwt.MapClaims) (uint64, *rest_err.RestErr) {
	if userID, ok := claims["id"].(float64); ok {
		return uint64(userID), nil
	}

	return 0, rest_err.NewBadRequestError("User ID not found in token")
}

package repository

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/users"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdatePasswordRepository(
	token string, userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdatePasswordRepository",
		zap.String("journey", "updatePassword"),
	)

	secretKey := env.GetEnv("SECRET_KEY")

	claims, err := parseToken(token, secretKey)

	if err != nil {
		logger.Error(
			"Error trying to validate token",
			err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	userID, err := extractUserID(claims)

	if err != nil {
		logger.Error(
			"Error trying to extract user id from token",
			err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	query := "UPDATE users SET password = $1 WHERE id = $2"

	statement, err := ur.db.Prepare(query)

	if err != nil {
		logger.Error(
			"Error preparing updatePassword statement",
			err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	if _, err := statement.Exec(userDomain.GetPassword(), userID); err != nil {
		logger.Error(
			"Error executing SQL statement",
			err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"UpdatePasswordRepository executed successfully",
		zap.String("journey", "updatePassword"),
	)

	return nil
}

func parseToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		logger.Error(
			"Error trying to parse token",
			err,
			zap.String("journey", "updatePassword"),
		)

		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	logger.Error(
		"Error trying to validate token",
		err,
		zap.String("journey", "updatePassword"),
	)

	return nil, err
}

func extractUserID(claims jwt.MapClaims) (uint64, error) {
	if userID, ok := claims["user_id"].(float64); ok {
		return uint64(userID), nil
	}

	return 0, rest_err.NewBadRequestError("User ID not found in token")
}

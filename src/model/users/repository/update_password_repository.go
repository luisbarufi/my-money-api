package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/configuration/utils"
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

	claims, rest_err := utils.ParseToken(token, secretKey)

	if rest_err != nil {
		logger.Error(
			"Error trying to validate token",
			rest_err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err
	}

	userID, rest_err := utils.ExtractUserID(claims)

	if rest_err != nil {
		logger.Error(
			"Error trying to extract user id from token",
			rest_err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err
	}

	query := "UPDATE users SET password = $1 WHERE id = $2"

	statement, err := ur.db.Prepare(query)

	if err != nil {
		logger.Error(
			"Error preparing updatePassword statement",
			err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err
	}

	defer statement.Close()

	if _, err := statement.Exec(userDomain.GetPassword(), userID); err != nil {
		logger.Error(
			"Error executing SQL statement",
			err,
			zap.String("journey", "updatePassword"),
		)

		return rest_err
	}

	logger.Info(
		"UpdatePasswordRepository executed successfully",
		zap.String("journey", "updatePassword"),
	)

	return nil
}

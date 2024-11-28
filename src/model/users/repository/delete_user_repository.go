package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUserRepository(userId uint64) *rest_err.RestErr {
	logger.Info("Init DeleteUserRepository", zap.String("journey", "deleteUser"))

	user, _ := ur.FindUserByIDRepository(userId)

	if user == nil {
		logger.Info("User ID not found", zap.String("journey", "deleteUser"))

		return rest_err.NewNotFoundError("User ID not found")
	}

	statement, err := ur.db.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		logger.Error(
			"Error preparing delete statement",
			err,
			zap.String("journey", "deleteUser"),
		)

		return nil
	}

	defer statement.Close()

	if _, err := statement.Exec(userId); err != nil {
		logger.Error(
			"Error executing delete statement",
			err,
			zap.String("journey", "deleteUser"),
		)

		return rest_err.NewInternalServerError("Error deleting user")
	}

	logger.Info(
		"DeleteUserRepository executed successfully",
		zap.String("journey", "deleteUser"),
	)

	return nil
}

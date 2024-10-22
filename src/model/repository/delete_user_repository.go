package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId uint64) *rest_err.RestErr {
	logger.Info(
		"Init updatedUser repository", zap.String("journey", "updatedUser"),
	)

	statement, err := ur.db.Conn.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return rest_err.NewInternalServerError("Error deleting user")
	}
	defer statement.Close()

	if _, err := statement.Exec(userId); err != nil {
		return rest_err.NewInternalServerError("Error deleting user")
	}

	logger.Info("User updated successfully",
		zap.String("journey", "updatedUser"),
	)
	return nil
}

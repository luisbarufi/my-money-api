package repository

import (
	"database/sql"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId uint64) *rest_err.RestErr {
	logger.Info("Init deleteUser repository", zap.String("journey", "deleteUser"))

	statement, restErr := ur.prepareDeleteUserStatement()
	if restErr != nil {
		return restErr
	}
	defer statement.Close()

	restErr = ur.executeDeleteUser(statement, userId)
	if restErr != nil {
		return restErr
	}

	logger.Info("User deleted successfully", zap.String("journey", "deleteUser"))
	return nil
}

func (ur *userRepository) prepareDeleteUserStatement() (
	*sql.Stmt, *rest_err.RestErr,
) {
	statement, err := ur.db.Conn.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		logger.Error("Error preparing delete statement",
			err,
			zap.String("journey", "deleteUser"),
		)
		return nil, rest_err.NewInternalServerError("Error deleting user")
	}
	return statement, nil
}

func (ur *userRepository) executeDeleteUser(
	statement *sql.Stmt, userId uint64,
) *rest_err.RestErr {
	if _, err := statement.Exec(userId); err != nil {
		logger.Error("Error executing delete statement",
			err,
			zap.String("journey", "deleteUser"),
		)
		return rest_err.NewInternalServerError("Error deleting user")
	}
	return nil
}

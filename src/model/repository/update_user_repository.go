package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId uint64, userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init updatedUser repository", zap.String("journey", "updatedUser"),
	)

	query, args := buildUpdateQuery(userId, userDomain)

	if len(args) == 1 {
		logger.Info(
			"No parameters were provided to update",
			zap.String("journey", "updatedUser"),
		)
		return nil
	}

	if err := ur.executeUpdateQuery(query, args); err != nil {
		return err
	}

	logger.Info("User updated successfully",
		// zap.String("userId", userId),
		zap.String("journey", "updatedUser"),
	)
	return nil
}

func buildUpdateQuery(userId uint64, userDomain model.UserDomainInterface) (
	string, []interface{},
) {
	value := converter.ConvertDomainToEntity(userDomain)
	updates := make([]string, 0)
	args := make([]interface{}, 0)

	fieldsToUpdate := map[string]interface{}{
		"name":       value.Name,
		"nick":       value.Nick,
		"updated_at": time.Now().UTC(),
	}

	for field, val := range fieldsToUpdate {
		if val != "" {
			updates = append(updates, fmt.Sprintf("%s = $%d", field, len(args)+1))
			args = append(args, val)
		}
	}

	args = append(args, userId)
	query := fmt.Sprintf(
		"UPDATE users SET %s WHERE id = $%d", strings.Join(updates, ", "), len(args),
	)

	return query, args
}

func (ur *userRepository) executeUpdateQuery(
	query string, args []interface{},
) *rest_err.RestErr {
	statement, err := ur.db.Conn.Prepare(query)
	if err != nil {
		logger.Error("Error preparing update user query", err)
		return rest_err.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	if _, err := statement.Exec(args...); err != nil {
		logger.Error("Error executing update user query", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}

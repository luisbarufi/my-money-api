package repository

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	query := buildSelectUserQuery("email")
	user, err := ur.executeFindUser(query, email)

	if err != nil || user.ID == 0 {
		logger.Error(
			"Error calling executeFindUser",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, err
	}

	logger.Info(
		"FindUserByEmail repository executed successfully",
		zap.String("id", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "findUserByEmail"),
	)
	return converter.ConvertEntityToDomain(*user), nil
}

func (ur *userRepository) FindUserByID(id uint64) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info(
		"Init findUserByID repository",
		zap.String("journey", "findUserByID"),
	)

	query := buildSelectUserQuery("id")
	user, err := ur.executeFindUser(query, id)
	if err != nil {
		logger.Error(
			"Error calling buildSelectUserQuery",
			err,
			zap.String("journey", "findUserByID"),
		)
		return nil, err
	}

	logger.Info(
		"FindUserByID repository executed successfully",
		zap.String("id", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "findUserByID"),
	)
	return converter.ConvertEntityToDomain(*user), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	query := buildSelectUserLoginQuery(password, email)
	user, err := ur.executeFindUserByEmailAndPassword(query)
	if err != nil {
		logger.Error(
			"Error calling executeFindUserByEmailAndPassword",
			err,
			zap.String("journey", "findUserByEmailAndPassword"),
		)
		return nil, err
	}

	logger.Info(
		"findUserByEmailAndPassword repository executed successfully",
		zap.String("id", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	return converter.ConvertEntityToDomain(*user), nil
}

func buildSelectUserQuery(field string) string {
	return fmt.Sprintf("SELECT * FROM users WHERE %s = $1", field)
}

func buildSelectUserLoginQuery(email, password string) string {
	return fmt.Sprintf(
		`SELECT * FROM users WHERE email = '%s' AND password = '%s'`,
		email,
		password,
	)
}

func (ur *userRepository) executeFindUser(query string, arg interface{}) (
	*entity.UserEntity, *rest_err.RestErr,
) {
	row, err := ur.db.Conn.Query(query, arg)
	if err != nil {
		logger.Error("Error executing find user query", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer row.Close()

	var user entity.UserEntity

	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			logger.Error("Error scanning find user result", err)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}
	return &user, nil
}

func (ur *userRepository) executeFindUserByEmailAndPassword(query string) (
	*entity.UserEntity, *rest_err.RestErr,
) {
	row, err := ur.db.Conn.Query(query)
	if err != nil {
		logger.Error("Error executing executeFindUserByEmailAndPassword query", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer row.Close()

	var user entity.UserEntity

	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			logger.Error("Error scanning find user result", err)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	return &user, nil
}

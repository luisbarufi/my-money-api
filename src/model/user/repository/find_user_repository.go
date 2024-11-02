package repository

import (
	"fmt"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/user"
	"github.com/luisbarufi/my-money-api/src/model/user/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/user/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmailRepository(email string) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	row, err := ur.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		logger.Error(
			"Error executing find user query",
			err,
			zap.String("journey", "findUserByEmail"),
		)
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
			logger.Error(
				"Error scanning find user result",
				err,
				zap.String("journey", "findUserByEmail"),
			)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	if user.ID == 0 {
		return nil, rest_err.NewNotFoundError("User not found")
	}

	logger.Info(
		"FindUserByEmail repository executed successfully",
		zap.String("id", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "findUserByEmail"),
	)
	return converter.ConvertEntityToDomain(user), nil
}

func (ur *userRepository) FindUserByIDRepository(id uint64) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info(
		"Init findUserByID repository",
		zap.String("journey", "findUserByID"),
	)

	row, err := ur.db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		logger.Error(
			"Error executing find user query",
			err,
			zap.String("journey", "findUserByID"),
		)
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
			logger.Error(
				"Error scanning find user result",
				err,
				zap.String("journey", "findUserByID"),
			)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	if user.ID == 0 {
		return nil, rest_err.NewNotFoundError("User not found")
	}

	logger.Info(
		"FindUserByID repository executed successfully",
		zap.String("id", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "findUserByID"),
	)
	return converter.ConvertEntityToDomain(user), nil
}

func (ur *userRepository) FindUserByEmailAndPasswordRepository(
	email, password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	query := "SELECT * FROM users WHERE email = $1 AND password = $2"
	row, err := ur.db.Query(query, email, password)
	if err != nil {
		logger.Error(
			"Error executing executeFindUserByEmailAndPassword query",
			err,
			zap.String("journey", "findUserByEmailAndPassword"),
		)
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
			logger.Error(
				"Error scanning find user result",
				err,
				zap.String("journey", "findUserByEmailAndPassword"),
			)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	if user.ID == 0 {
		return nil, rest_err.NewNotFoundError("Email or password incorrect")
	}

	logger.Info(
		"findUserByEmailAndPassword repository executed successfully",
		zap.String("id", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	return converter.ConvertEntityToDomain(user), nil
}

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

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))

	query, args := buildInsertUserQuery(userDomain)

	user, err := ur.executeInsertUser(query, args)
	if err != nil {
		return nil, err
	}

	logger.Info("CreateUser repository executed successfully",
		zap.String("userId", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(*user), nil
}

func buildInsertUserQuery(userDomain model.UserDomainInterface) (
	string, []interface{},
) {
	query := `INSERT INTO users (name, nick, email, password)
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id, name, nick, email, password, created_at, updated_at`
	value := converter.ConvertDomainToEntity(userDomain)

	args := []interface{}{
		value.Name,
		value.Nick,
		value.Email,
		value.Password,
	}

	return query, args
}

func (ur *userRepository) executeInsertUser(query string, args []interface{}) (
	*entity.UserEntity, *rest_err.RestErr,
) {
	row, err := ur.db.Conn.Query(query, args...)
	if err != nil {
		logger.Error("Error executing insert user query", err)
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
			logger.Error("Error scanning insert user result", err)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	return &user, nil
}

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

func (ur *userRepository) CreateUserRepository(
	userDomain model.UserDomainInterface) (
	model.UserDomainInterface, *rest_err.RestErr,
) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))

	query := `INSERT INTO users (name, nick, email, password)
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id, name, nick, email, password, created_at, updated_at`

	value := converter.ConvertDomainToEntity(userDomain)

	row, err := ur.db.Conn.Query(
		query,
		value.Name,
		value.Nick,
		value.Email,
		value.Password,
	)
	if err != nil {
		logger.Error(
			"Error executing insert user query",
			err,
			zap.String("journey", "createUser"),
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
				"Error scanning insert user result",
				err,
				zap.String("journey", "createUser"),
			)
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	logger.Info("CreateUser repository executed successfully",
		zap.String("userId", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(user), nil
}

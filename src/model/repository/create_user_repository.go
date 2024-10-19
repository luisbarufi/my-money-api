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

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))

	query := `
		INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, email, password, created_at, updated_at`

	value := converter.ConvertDomainToEntity(userDomain)

	row, err := ur.db.Conn.Query(query, value.Name, value.Email, value.Password)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer row.Close()

	var user entity.UserEntity

	if row.Next() {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, rest_err.NewInternalServerError("Erro ao escanear os resultados")
		}
	} else {
		return nil, rest_err.NewInternalServerError("Nenhum resultado retornado")
	}

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("UserId", fmt.Sprintf("%d", user.ID)),
		zap.String("journey", "createUser"))

	return converter.ConvertEntityToDomain(user), nil
}

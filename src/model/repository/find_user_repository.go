package repository

import (
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity"
	"github.com/luisbarufi/my-money-api/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository", zap.String("journey", "findUserByID"))

	row, err := ur.db.Conn.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, rest_err.NewNotFoundError("Usuário não encontrado")
	}
	defer row.Close()

	var user entity.UserEntity

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			logger.Error("Error scanning results", err, zap.String("journey", "createUser"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	return converter.ConvertEntityToDomain(user), nil
}

func (ur *userRepository) FindUserByID(id uint64) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository", zap.String("journey", "findUserByID"))

	row, err := ur.db.Conn.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, rest_err.NewNotFoundError("Usuário não encontrado")
	}
	defer row.Close()

	var user entity.UserEntity

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			logger.Error("Error scanning results", err, zap.String("journey", "createUser"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	return converter.ConvertEntityToDomain(user), nil
}

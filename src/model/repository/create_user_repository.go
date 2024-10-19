package repository

import (
	"time"

	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	query := `
		INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	var userID int64
	var createdAt, updatedAt time.Time

	err := ur.db.Conn.QueryRow(
		query, userDomain.GetName(),
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	).Scan(&userID, &createdAt, &updatedAt)

	if err != nil {
		logger.Error("Error inserting user into database", err)
		return nil, rest_err.NewInternalServerError("Error inserting user into database")
	}

	userDomain.SetID(userID)
	userDomain.SetCreatedAt(createdAt)
	userDomain.SetUpdatedAt(updatedAt)

	logger.Info("User inserted successfully!")

	return userDomain, nil
}

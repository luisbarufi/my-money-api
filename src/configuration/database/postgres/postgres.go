package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
)

func InitConnection() {
	connStr := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
		env.GetEnv("DATABASE_USER"),
		env.GetEnv("DATABASE_PASSWORD"),
		env.GetEnv("DATABASE_HOST"),
		env.GetEnv("DATABASE_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logger.Info("Postgres connection established!")
}

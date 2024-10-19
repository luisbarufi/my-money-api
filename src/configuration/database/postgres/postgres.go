package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
)

var (
	CONN_STR = "CONN_STR"
)

type Database struct {
	Conn *sql.DB
}

func NewPostgresConnection() (*Database, error) {
	connStr := env.GetEnv(CONN_STR)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logger.Info("PostgreSQL connection established successfully!")

	return &Database{Conn: db}, nil
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		logger.Error("Error closing database connection!", err)
	} else {
		logger.Info("Database connection closed!")
	}
}

package postgres

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
)

var (
	CONN_STR   = "CONN_STR"
	SOURCE_URL = "SOURCE_URL"
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

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		env.GetEnv(SOURCE_URL),
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	logger.Info("PostgreSQL connection established successfully!")
	logger.Info("Migrations successfully implemented!")

	return &Database{Conn: db}, nil
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		logger.Error("Error closing database connection!", err)
	} else {
		logger.Info("Database connection closed!")
	}
}

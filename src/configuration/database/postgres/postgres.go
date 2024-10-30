package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
)

var (
	SOURCE_URL  = "SOURCE_URL"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
)

func NewPostgresConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.GetEnv(DB_HOST),
		env.GetEnv(DB_PORT),
		env.GetEnv(DB_USER),
		env.GetEnv(DB_PASSWORD),
		env.GetEnv(DB_NAME),
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
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

	err = db.Ping()
	return db, err
}

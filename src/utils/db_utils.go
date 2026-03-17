package utils

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func ConnectToDatabase() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", "user:password@localhost/database")
}
package config

import (
	"database/sql"
	_"github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {

	database, err := sql.Open("postgres", "postgres://postgres:root@localhost/labora_proyect_1?sslmode=disable")

	return database, err
}

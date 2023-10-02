// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "database"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db: db,
	}
}

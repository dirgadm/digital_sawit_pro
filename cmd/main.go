package main

import (
	"database/sql"
	"fmt"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := "postgres://postgres:postgres@db:5432/database?sslmode=disable"

	fmt.Println("===================================", dbDsn)
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}

	fmt.Println("==============123=====================", opts)

	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@db:5432/database?sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Println("==============123=====================", db)
	return handler.NewServer(opts)
}

package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ranggabudipangestu/hexagonal-api/database"
	"github.com/ranggabudipangestu/hexagonal-api/database/migration"
	"github.com/ranggabudipangestu/hexagonal-api/internal/factory"
	"github.com/ranggabudipangestu/hexagonal-api/internal/http"
	"github.com/ranggabudipangestu/hexagonal-api/internal/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database.GetConnection()
}

func main() {
	database.CreateConnection()

	var m string // for check migration

	flag.StringVar(
		&m,
		"m",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration
to use this flag:
	use -m=migrate for migrate table
	use -m=rollback for rollback table
	use -m=status for get status migration`,
	)

	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	} else if m == "rollback" {
		migration.Rollback()
	} else if m == "status" {
		migration.Status()
	}

	f := factory.NewFactory()
	e := echo.New()

	middleware.LogMiddleware(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}

package main

import (
	"github.com/ranggabudipangestu/mvc-api/config"
	"github.com/ranggabudipangestu/mvc-api/middlewares"
	"github.com/ranggabudipangestu/mvc-api/routes"
)

func main() {

	config.InitDB()
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":5000"))

}

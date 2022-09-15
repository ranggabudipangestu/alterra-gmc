package main

import (
	"github.com/ranggabudipangestu/mvc-api/config"
	"github.com/ranggabudipangestu/mvc-api/routes"
)

func main() {

	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":5000"))

}

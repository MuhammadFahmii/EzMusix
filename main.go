package main

import (
	"EzMusix/config"
	"EzMusix/routes"
)

func main() {
	config.InitDB()
	e := routes.NewRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}

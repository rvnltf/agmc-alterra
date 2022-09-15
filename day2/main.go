package main

import (
	"agmc-api/config"
	"agmc-api/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":3000"))
}

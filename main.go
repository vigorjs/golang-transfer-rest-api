package main

import (
	"golang-rest-api/initializers"
	"golang-rest-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Conn()
	initializers.MigrateDb()
}

func main() {
	routes.SetupRouter();
}

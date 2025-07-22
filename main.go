package main

import (
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/databases/migrations"
	"aplikasi-adakost-be/routers"
)

// @title AdaKost API
// @version 1.0
// @description API dokumentasi untuk Aplikasi AdaKost
// @BasePath /api
func main() {
	connection.DbConnection()
	migrations.DBMigrate(connection.DBConnections)
	r := routers.SetupRouters()
	r.Run(":8080")
	// Init()
}

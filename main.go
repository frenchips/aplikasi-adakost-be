package main

import (
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/databases/migrations"
)

func main() {
	connection.DbConnection()
	migrations.DBMigrate(connection.DBConnections)
	// Init()
}

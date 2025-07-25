package main

import (
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/databases/migrations"
	"aplikasi-adakost-be/routers"
	"os"
)

// @title Aplikasi Booking Kost API
// @BasePath /api
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer " followed by your JWT token.
// @Security BearerAuth
func main() {
	connection.DbConnection()
	migrations.DBMigrate(connection.DBConnections)
	r := routers.SetupRouters()

	// Ambil port dari env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback untuk lokal
	}

	r.Run(":" + port)
	// Init()
}

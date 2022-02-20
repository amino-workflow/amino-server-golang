package main

import (
	"amino-server/config"
	"amino-server/migrations"
)

func main() {
	// Load config
	config.Load()

	// Start DB
	config.StartDB()
	defer config.CloseDB()

	// Run Migration
	migrations.Migrate()

	// Start Server
	// r := gin.Default()
	// port := config.Get("SERVER_PORT")
	// r.Run(port)
}

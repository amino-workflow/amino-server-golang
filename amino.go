package main

import (
	"amino-server/config"
	"amino-server/migrations"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.Load()

	// Start DB
	config.StartDB()
	defer config.CloseDB()

	// Run Migration
	fmt.Println("Running Migrations...")
	migrations.Migrate()
	fmt.Println("Migrations Complete!")

	fmt.Println("Starting Server...")
	// Start Server
	r := gin.Default()
	port := config.Get("SERVER_PORT")
	r.Run(port)
	defer fmt.Println("Shuting Down Server...")
}

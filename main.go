package main

import (
	"toko-buah/config/db"
	"toko-buah/routes"
)

func main() {

	// Server to database
	db.Server()

	// Initalize the router
	routes.SetupRouter()

	// Run the server
	routes.SetupRouter().Run(":3000")
}

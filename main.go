package main

import (
	"toko-buah/config/db"
	"toko-buah/routes"
)

// func main() {

// 	// Server to database
// 	db.Server()

// 	// Initalize the router
// 	routes.SetupRouter()

// 	// Run the server
// 	routes.SetupRouter().Run(":3000")
// }

func main() {
	// Server to database
	db.Server()

	// Inisialisasi router
	router := routes.SetupRouter()

	// Tentukan port yang akan digunakan
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Port default jika tidak ada variabel lingkungan PORT
	}

	// Jalankan server dengan menggunakan HTTP
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

package main

import (
	"fmt"
	"net/http"
	"anjalisreekumar.github/book-manager/routes"
)

func main() {
	// Register the routes
	routes.RegisterRoutes()

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

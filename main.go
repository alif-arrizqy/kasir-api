package main

import (
	"fmt"
	"kasir-api/routes"
	"net/http"
)

func main() {
	// setup routes
	routes.SetupRoutes()

	fmt.Println("Server is running on port 8889")

	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}

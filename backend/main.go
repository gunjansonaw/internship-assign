package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"backend/routes"
)

func main() {
	// Set up routes
	routes.SetHolidayRoutes()

	// Enable CORS
	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS()(http.DefaultServeMux)))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"backend/routes"
)

func main() {
	// Initialize a new mux router
	router := mux.NewRouter()

	// Set up routes
	routes.SetHolidayRoutes(router)

	// Define allowed CORS settings
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
	)(router)

	// Start the server with CORS enabled
	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", corsHandler))
}

package main

import (
	"log"
	"net/http"

	"backend/routes"
	"backend/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	utils.ConnectDB()

	// Initialize a new mux router
	router := mux.NewRouter()

	// Set up routes
	routes.SetHolidayRoutes(router)

	// Define allowed CORS settings
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	// Start the server
	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", corsHandler))
}

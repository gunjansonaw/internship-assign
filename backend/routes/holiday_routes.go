package routes

import (
	"backend/handlers"
	"github.com/gorilla/mux"
)

func SetHolidayRoutes(router *mux.Router) {
	router.HandleFunc("/api/holiday", handlers.GetHolidays).Methods("GET")       // Get all holidays
	router.HandleFunc("/api/holiday", handlers.AddHoliday).Methods("POST")       // Add a new holiday
	router.HandleFunc("/api/holiday/{id:[a-fA-F0-9]{24}}", handlers.DeleteHoliday).Methods("DELETE") // Delete a holiday by ID
}

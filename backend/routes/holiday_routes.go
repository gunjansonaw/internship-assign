package routes

import (
	"backend/handlers"
	"net/http"
)

func SetHolidayRoutes() {
	// Define the routes
	http.HandleFunc("/api/holiday", handlers.GetHolidays) // Get all holidays and add holiday
	http.HandleFunc("/api/holidays/:id", handlers.DeleteHoliday) // Delete holiday
	http.HandleFunc("/api/holiday/",handlers.AddHoliday)
}

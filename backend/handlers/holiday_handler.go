package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Holiday struct {
	ID   int    `json:"id"`
	Date int    `json:"date"`
	Name string `json:"name"`
	Year int    `json:"year"`
	Month int   `json:"month"`
}

// In-memory data store
var holidays = make(map[int]Holiday)
var nextID = 1

// Handle GET request to fetch holidays by month and year
func GetHolidays(w http.ResponseWriter, r *http.Request) {
	// Extract month and year from query parameters
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")
	if month == "" || year == "" {
		http.Error(w, "Month and Year are required", http.StatusBadRequest)
		return
	}

	// Parse month and year as integers
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	// Collect holidays for the specified month and year
	var result []Holiday
	for _, holiday := range holidays {
		if holiday.Month == monthInt && holiday.Year == yearInt {
			result = append(result, holiday)
		}
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Handle POST request to add a holiday
func AddHoliday(w http.ResponseWriter, r *http.Request) {
	// Parse the holiday data from the request body
	var holiday Holiday
	err := json.NewDecoder(r.Body).Decode(&holiday)
	if err != nil {
		http.Error(w, "Invalid holiday data", http.StatusBadRequest)
		return
	}

	// Assign an ID and store the holiday
	holiday.ID = nextID
	nextID++
	holidays[holiday.ID] = holiday

	// Respond with the added holiday
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(holiday)
}

func DeleteHoliday(w http.ResponseWriter, r *http.Request) {
    // Extract the holiday ID from the URL path variables
    vars := mux.Vars(r)
    idStr, ok := vars["id"]
    if !ok {
        http.Error(w, "Holiday ID is required", http.StatusBadRequest)
        return
    }

    // Convert the ID to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
        return
    }

    // Delete the holiday if it exists
    if _, exists := holidays[id]; !exists {
        http.Error(w, "Holiday not found", http.StatusNotFound)
        return
    }
    delete(holidays, id)

    // Send a success response
    w.WriteHeader(http.StatusNoContent)
}
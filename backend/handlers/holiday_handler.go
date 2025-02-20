package handlers

import (
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handle GET request to fetch holidays by month and year
func GetHolidays(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	if month == "" || year == "" {
		http.Error(w, "Month and Year are required", http.StatusBadRequest)
		return
	}

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

	collection := utils.GetCollection("holidays")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"month": monthInt, "year": yearInt}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	var holidays []models.Holiday
	if err := cursor.All(ctx, &holidays); err != nil {
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(holidays)
}

// Handle POST request to add a holiday
func AddHoliday(w http.ResponseWriter, r *http.Request) {
	var holiday models.Holiday
	err := json.NewDecoder(r.Body).Decode(&holiday)
	if err != nil {
		http.Error(w, "Invalid holiday data", http.StatusBadRequest)
		return
	}

	holiday.ID = primitive.NewObjectID()

	collection := utils.GetCollection("holidays")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, holiday)
	if err != nil {
		http.Error(w, "Failed to insert holiday", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(holiday)
}

// Handle DELETE request to remove a holiday
func DeleteHoliday(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Holiday ID is required", http.StatusBadRequest)
		return
	}

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	collection := utils.GetCollection("holidays")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Holiday not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

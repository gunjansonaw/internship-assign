package controllers

import (
	"context"
	"net/http"
	"time"

	"holiday-calendar/config"
	"holiday-calendar/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var holidayCollection = config.GetCollection("holidays")

// Add a Holiday
func AddHoliday(c *gin.Context) {
	var holiday models.Holiday
	if err := c.ShouldBindJSON(&holiday); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	holiday.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := holidayCollection.InsertOne(ctx, holiday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add holiday"})
		return
	}

	c.JSON(http.StatusCreated, holiday)
}

// Get all Holidays
func GetHolidays(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := holidayCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch holidays"})
		return
	}

	var holidays []models.Holiday
	if err = cursor.All(ctx, &holidays); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading holidays"})
		return
	}

	c.JSON(http.StatusOK, holidays)
}

// Delete a Holiday
func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = holidayCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete holiday"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Holiday deleted"})
}

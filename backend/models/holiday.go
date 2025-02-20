package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Holiday struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Date  int                `bson:"date" json:"date"` // Date in the month
	Month int                `bson:"month" json:"month"`
	Year  int                `bson:"year" json:"year"`
}

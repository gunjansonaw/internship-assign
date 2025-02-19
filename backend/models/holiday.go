package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Holiday struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Date  string             `bson:"date" json:"date"`
	Month int                `bson:"month" json:"month"`
	Year  int                `bson:"year" json:"year"`
}

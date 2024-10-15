package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Description string             `bson:"description" json:"description" validate:"required"`
	Genre       string             `bson:"genre" json:"genre" validate:"required"`
	Duration    string                `bson:"duration" json:"duration" validate:"required"`
	ReleaseDate string             `bson:"release_date" json:"release_date" validate:"required"`
	Rating      string            `bson:"rating" json:"rating" validate:"required"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
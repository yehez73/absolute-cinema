package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Theater struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Location  string             `bson:"location" json:"location" validate:"required"`
	MaxRows   string             `bson:"max_rows" json:"max_rows" validate:"required"`
	MaxCols   int                `bson:"max_cols" json:"max_cols" validate:"required"`
	Seats     []SeatInfo         `bson:"seats" json:"seats"`          // Array of seats
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type SeatInfo struct {
	SeatCode    string `bson:"seat_code" json:"seat_code"`
	IsAvailable bool   `bson:"is_available" json:"is_available"`
}

type Seat struct {
	Row string `bson:"row" json:"row"`
	Col int    `bson:"col" json:"col"`
}
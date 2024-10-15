package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	ShowtimeID  primitive.ObjectID `bson:"showtime_id" json:"showtime_id" validate:"required"`
	Seats       []string           `bson:"seat_number" json:"seat_number" validate:"required"`
	TotalPrice  int                `bson:"total_price" json:"total_price"`
	BookingDate string             `bson:"booking_date" json:"booking_date"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Showtime struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MovieID        primitive.ObjectID `bson:"movie_id" json:"movie_id" validate:"required"`
	TheaterID      primitive.ObjectID `bson:"theater_id" json:"theater_id" validate:"required"`
	ShowDate       string             `bson:"show_date" json:"show_date" validate:"required"`
	StartTime      string             `bson:"start_time" json:"start_time" validate:"required"`
	EndTime        string             `bson:"end_time" json:"end_time" validate:"required"`
	AvailableSeats []SeatInfo         `bson:"available_seats" json:"available_seats"`
	Price          int                `bson:"pricing" json:"pricing" validate:"required"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type TheaterShowtime struct {
	TheaterID      primitive.ObjectID `json:"theater_id"`
	Name   		   string             `json:"theater_name"`
	Location 	   string            `json:"theater_location"`
	Showtimes 	   []Showtime         `json:"showtimes"`
}

type GroupedShowtime struct {
	Theaters	   []TheaterShowtime `json:"theaters"`
}
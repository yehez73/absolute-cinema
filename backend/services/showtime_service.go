package services

import (
	"backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ShowtimeCollection *mongo.Collection

func InitShowtimeService(db *mongo.Database) {
	ShowtimeCollection = db.Collection("showtime")
}

func GetShowtimes() ([]models.Showtime, error) {
	var showtimes []models.Showtime

	cursor, err := ShowtimeCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return showtimes, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var showtime models.Showtime
		cursor.Decode(&showtime)
		showtimes = append(showtimes, showtime)
	}

	return showtimes, nil
}

func GetSpecShowtime(id string) (models.Showtime, error) {
	var showtime models.Showtime
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return showtime, err
	}

	err = ShowtimeCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&showtime)
	return showtime, err
}

func CreateShowtime(showtime *models.Showtime) error {
	movieCollection := ShowtimeCollection.Database().Collection("movies")
	var movie models.Movie
	err := movieCollection.FindOne(context.Background(), bson.M{"_id": showtime.MovieID}).Decode(&movie)
	if err != nil {
		return err
	}

	var theater models.Theater
	theaterCollection := ShowtimeCollection.Database().Collection("theater")
	err = theaterCollection.FindOne(context.Background(), bson.M{"_id": showtime.TheaterID}).Decode(&theater)
	if err != nil {
		return err
	}

	// Fetch available seats of the theater
	var availableSeats []models.SeatInfo
	for _, seat := range theater.Seats {
		availableSeats = append(availableSeats, models.SeatInfo{
			SeatCode:    seat.SeatCode,
			IsAvailable: true,
		})
	}

	showtime.AvailableSeats = availableSeats
	showtime.CreatedAt = time.Now()

	_, err = ShowtimeCollection.InsertOne(context.Background(), showtime)
	return err
}

func UpdateShowtime(id string, showtime *models.Showtime) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	movieCollection := ShowtimeCollection.Database().Collection("movies")
	var movie models.Movie
	err = movieCollection.FindOne(context.Background(), bson.M{"_id": showtime.MovieID}).Decode(&movie)
	if err != nil {
		return err
	}

	var theater models.Theater
	theaterCollection := ShowtimeCollection.Database().Collection("theater")
	err = theaterCollection.FindOne(context.Background(), bson.M{"_id": showtime.TheaterID}).Decode(&theater)
	if err != nil {
		return err
	}

	// Fetch available seats of the theater
	var availableSeats []models.SeatInfo
	for _, seat := range theater.Seats {
		availableSeats = append(availableSeats, models.SeatInfo{
			SeatCode:    seat.SeatCode,
			IsAvailable: true,
		})
	}

	var existingShowtime models.Showtime
	err = ShowtimeCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&existingShowtime)
	if err != nil {
		return err
	}

	showtime.AvailableSeats = availableSeats
	showtime.CreatedAt = existingShowtime.CreatedAt
	showtime.UpdatedAt = time.Now()

	_, err = ShowtimeCollection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, showtime)
	return err
}

func DeleteShowtime(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ShowtimeCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
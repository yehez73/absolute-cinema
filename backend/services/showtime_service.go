package services

import (
	"backend/models"
	"context"
	"sort"
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

func GetSpecShowtimeByMovieDate(movieID string, showdate string) (models.GroupedShowtime, error) {
	var showtimes []models.Showtime
	var groupedResponse models.GroupedShowtime

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return groupedResponse, err
	}

	cursor, err := ShowtimeCollection.Find(context.Background(), bson.M{"movie_id": objectID, "show_date": showdate})
	if err != nil {
		return groupedResponse, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &showtimes); err != nil {
		return groupedResponse, err
	}

	sort.Slice(showtimes, func(i, j int) bool {
		startTimeI, _ := time.Parse("15:04", showtimes[i].StartTime)
		startTimeJ, _ := time.Parse("15:04", showtimes[j].StartTime)
		return startTimeI.Before(startTimeJ)
	})

	theaterMap := make(map[string][]models.Showtime)
	for _, showtime := range showtimes {
		theaterID := showtime.TheaterID.Hex()
		theaterMap[theaterID] = append(theaterMap[theaterID], showtime)
	}

	theaterCollection := ShowtimeCollection.Database().Collection("theater")
	for theaterID, shows := range theaterMap {
		theaterObjectID, err := primitive.ObjectIDFromHex(theaterID)
		if err != nil {
			return groupedResponse, err
		}

		var theater models.Theater
		err = theaterCollection.FindOne(context.Background(), bson.M{"_id": theaterObjectID}).Decode(&theater)
		if err != nil {
			return groupedResponse, err
		}

		groupedResponse.Theaters = append(groupedResponse.Theaters, models.TheaterShowtime{
			TheaterID: theaterObjectID,
			Name:      theater.Name,
			Location:  theater.Location,
			Showtimes: shows,
		})
	}

	sort.Slice(groupedResponse.Theaters, func(i, j int) bool {
		return groupedResponse.Theaters[i].Name < groupedResponse.Theaters[j].Name
	})

	return groupedResponse, nil
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
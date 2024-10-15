package services

import (
	"backend/models"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var theaterCollection *mongo.Collection

func InitTheaterService(db *mongo.Database) {
	theaterCollection = db.Collection("theater")
}

func GetTheaters() ([]models.Theater, error) {
	var theaters []models.Theater
	cursor, err := theaterCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return theaters, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var theater models.Theater
		cursor.Decode(&theater)
		theaters = append(theaters, theater)
	}

	return theaters, nil
}

func GetSpecTheater(id string) (models.Theater, error) {
	var theater models.Theater
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return theater, err
	}

	err = theaterCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&theater)
	return theater, err
}

func CreateTheater(theater *models.Theater) error {
	var seats []models.Seat
	// Generate seats based on max rows and columns
	for r := 'A'; r <= rune(theater.MaxRows[0]); r++ {
		for c := 1; c <= theater.MaxCols; c++ {
			seats = append(seats, models.Seat{
				Row: string(r),
				Col: c,
			})
		}
	}
	
	var seatInfos []models.SeatInfo
	for _, seat := range seats {
		seatInfos = append(seatInfos, models.SeatInfo{
			SeatCode:    seat.Row + strconv.Itoa(seat.Col),
			IsAvailable: true,	// Seats Availabie by default
		})
	}
	theater.Seats = seatInfos
	theater.CreatedAt = time.Now()

	_, err := theaterCollection.InsertOne(context.Background(), theater)
	return err
}

func UpdateTheater(id string, theater *models.Theater) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var seats []models.Seat
	for r := 'A'; r <= rune(theater.MaxRows[0]); r++ {
		for c := 1; c <= theater.MaxCols; c++ {
			seats = append(seats, models.Seat{
				Row: string(r),
				Col: c,
			})
		}
	}

	var seatInfos []models.SeatInfo
	for _, seat := range seats {
		seatInfos = append(seatInfos, models.SeatInfo{
			SeatCode:    seat.Row + strconv.Itoa(seat.Col),
			IsAvailable: true,	// Seats Availabie by default
		})
	}
	theater.Seats = seatInfos

	var existingTheater models.Theater
	err = theaterCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&existingTheater)
	if err != nil {
		return err
	}

	theater.CreatedAt = existingTheater.CreatedAt
	theater.UpdatedAt = time.Now()

	_, err = theaterCollection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": theater})
	return err
}

func DeleteTheater(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = theaterCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
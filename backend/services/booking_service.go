package services

import (
	"backend/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookingCollection *mongo.Collection

func InitBookingService(db *mongo.Database) {
	bookingCollection = db.Collection("booking")
}

func GetBookings() ([]models.Booking, error) {
	var bookings []models.Booking

	cursor, err := bookingCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var booking models.Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func GetSpecBooking(id string) (models.Booking, error) {
	var booking models.Booking
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return booking, err
	}

	err = bookingCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&booking)
	return booking, err
}

func CreateBooking(booking *models.Booking, ID string) error {
	booking.UserID, _ = primitive.ObjectIDFromHex(ID)

	showtimeCollection := bookingCollection.Database().Collection("showtime")
	var showtime models.Showtime
	err := showtimeCollection.FindOne(context.Background(), bson.M{"_id": booking.ShowtimeID}).Decode(&showtime)
	if err != nil {
		return err
	}

	for _, seat := range booking.Seats {
		var seatStatus models.SeatInfo
		err = showtimeCollection.FindOne(
			context.Background(), 
			bson.M{"_id": showtime.ID, "available_seats.seat_code": seat, "available_seats.is_available": false}).Decode(&seatStatus)
		if err == nil {
			return errors.New("seat " + seat + " is already booked")
		}
	}

	booking.TotalPrice = showtime.Price * len(booking.Seats)
	booking.BookingDate = showtime.ShowDate + " " + showtime.StartTime
	booking.CreatedAt = time.Now()
	booking.Status = "Unpaid"

	for _, seat := range booking.Seats {
		_, err := showtimeCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": showtime.ID, "available_seats.seat_code": seat},
			bson.M{"$set": bson.M{"available_seats.$.is_available": false}},
		)
		if err != nil {
			return err
		}
	}

	_, err = bookingCollection.InsertOne(context.Background(), booking)
	return err
}

func UpdateBooking(id string, booking *models.Booking) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var existingBooking models.Booking
	err = bookingCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&existingBooking)
	if err != nil {
		return err
	}

	showtimeCollection := bookingCollection.Database().Collection("showtime")
	var showtime models.Showtime
	err = showtimeCollection.FindOne(context.Background(), bson.M{"_id": booking.ShowtimeID}).Decode(&showtime)
	if err != nil {
		return err
	}

	for _, seat := range existingBooking.Seats {
		if !func(slice []string, item string) bool {
			for _, s := range slice {
				if s == item {
					return true
				}
			}
			return false
		}(booking.Seats, seat) {
			_, err := showtimeCollection.UpdateOne(
				context.Background(),
				bson.M{"_id": showtime.ID, "available_seats.seat_code": seat},
				bson.M{"$set": bson.M{"available_seats.$.is_available": true}},
			)
			if err != nil {
				return err
			}
		}
	}

	for _, seat := range booking.Seats {
		if !func(slice []string, item string) bool {
			for _, s := range slice {
				if s == item {
					return true
				}
			}
			return false
		}(existingBooking.Seats, seat) {
			_, err := showtimeCollection.UpdateOne(
				context.Background(),
				bson.M{"_id": showtime.ID, "available_seats.seat_code": seat},
				bson.M{"$set": bson.M{"available_seats.$.is_available": false}},
			)
			if err != nil {
				return err
			}
		}
	}

	booking.TotalPrice = showtime.Price * len(booking.Seats)
	booking.BookingDate = showtime.ShowDate + " " + showtime.StartTime
	booking.Status = "Unpaid"
	booking.CreatedAt = existingBooking.CreatedAt
	booking.UpdatedAt = time.Now()

	_, err = bookingCollection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, booking)
	return err
}

func DeleteBooking(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var existingBooking models.Booking
	err = bookingCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&existingBooking)
	if err != nil {
		return err
	}

	showtimeCollection := bookingCollection.Database().Collection("showtime")
	var showtime models.Showtime
	err = showtimeCollection.FindOne(context.Background(), bson.M{"_id": existingBooking.ShowtimeID}).Decode(&showtime)
	if err != nil {
		return err
	}

	for _, seat := range existingBooking.Seats {
		_, err := showtimeCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": showtime.ID, "available_seats.seat_code": seat},
			bson.M{"$set": bson.M{"available_seats.$.is_available": true}},
		)
		if err != nil {
			return err
		}
	}

	_, err = bookingCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}

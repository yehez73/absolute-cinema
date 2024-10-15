package services

import (
	"backend/models"
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection

func InitMovieService(db *mongo.Database) {
	movieCollection = db.Collection("movies")
}

func GetMovies() ([]models.Movie, error) {
	var movies []models.Movie

	cursor, err := movieCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return movies, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	return movies, nil
}

func GetSpecMovie(id string) (models.Movie, error) {
	var movie models.Movie
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return movie, err
	}

	err = movieCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&movie)
	return movie, err
}

func CreateMovie(movie *models.Movie) error {
	movie.CreatedAt = time.Now()

	durationMinutes, err := strconv.Atoi(movie.Duration)
	if err != nil {
		return err
	}
	hours := durationMinutes / 60
	minutes := durationMinutes % 60
	movie.Duration = fmt.Sprintf("%dh %dm", hours, minutes)

	_, err = movieCollection.InsertOne(context.Background(), movie)
	return err
}

func UpdateMovie(id string, movie *models.Movie) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	durationMinutes, err := strconv.Atoi(movie.Duration)
	if err != nil {
		return err
	}
	hours := durationMinutes / 60
	minutes := durationMinutes % 60
	movie.Duration = fmt.Sprintf("%dh %dm", hours, minutes)

	var existingMovie models.Movie
	err = movieCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&existingMovie)
	if err != nil {
		return err
	}

	movie.CreatedAt = existingMovie.CreatedAt
	movie.UpdatedAt = time.Now()

	_, err = movieCollection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, movie)
	return err
}

func DeleteMovie(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = movieCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
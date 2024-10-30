package services

import (
	"backend/models"
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/spf13/viper"
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

func GetNowShowing() ([]models.Movie, error) {
	var movies []models.Movie

	cursor, err := movieCollection.Find(context.Background(), bson.M{"release_date": bson.M{"$lte": time.Now().Format("2006-01-02")}})

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

func GetUpcoming() ([]models.Movie, error) {
	var movies []models.Movie

	cursor, err := movieCollection.Find(context.Background(), bson.M{"release_date": bson.M{"$gt": time.Now().Format("2006-01-02")}})

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

func CreateMovie(movie *models.Movie, file []byte) error {
	account := viper.GetString("R2_ACCOUNT_ID")
	accessKey := viper.GetString("R2_ACCESS_KEY")
	secretKey := viper.GetString("R2_SECRET_KEY")
	bucket := viper.GetString("R2_BUCKET_NAME")

	r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", account)
	
	durationMinutes, err := strconv.Atoi(movie.Duration)
	if err != nil {
		log.Println(err)
		return err
	}
	hours := durationMinutes / 60
	minutes := durationMinutes % 60

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("auto"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: r2Endpoint,
			}, nil
		})),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	s3Client := s3.NewFromConfig(cfg)

	filename := uuid.New().String()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(filename),
		Body: bytes.NewReader(file),
		ContentType: aws.String("image/jpeg"),
	}

	_, err = s3Client.PutObject(context.TODO(), input)
	if err != nil {
		log.Println(err)
		return err
	}

	movie.Image = fmt.Sprintf("https://cineplex.image-assets.workers.dev/%s", filename)
	movie.Duration = fmt.Sprintf("%dh %dm", hours, minutes)
	movie.CreatedAt = time.Now()

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
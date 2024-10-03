package services

import (
	"backend/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetUsers() ([]models.User, error) { // For all users
	var users []models.User

	// This feels wrong...
	cursor, err := userCollection.Find(context.Background(), bson.M{}, options.Find().SetProjection(bson.M{"password": 0}))

	if err != nil {
		return users, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id string) (models.User, error) { // Just one user
	var user models.User
	objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return user, err
    }

	err = userCollection.FindOne(context.Background(), bson.M{"_id": objectID}, options.FindOne().SetProjection(bson.M{"password": 0})).Decode(&user)
    return user, err
}

func CreateUser(user models.User) error {
	if len(user.Password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }

	if user.Role != "admin" && user.Role != "member" {
        return errors.New("role not valid")
    }
    
    hashedPassword := hashPassword(user.Password)
    user.Password = hashedPassword

    user.CreatedAt = time.Now()

	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func UpdateUser(id string, user models.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	_, err = userCollection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": user})
	return err
}

func DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = userCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// This might be wrong...
var userCollection *mongo.Collection

func InitUserService(db *mongo.Database) {
    userCollection = db.Collection("users")
}

func Register(user models.User) error {
    if len(user.Password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }
    
    hashedPassword := hashPassword(user.Password)
    user.Password = hashedPassword

    user.CreatedAt = time.Now()

    if user.Role == "" {
        user.Role = "member"
    }

    _, err := userCollection.InsertOne(context.Background(), user)
    return err
}

func Login(email, password string) (models.User, error) {
    var user models.User
    err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return user, errors.New("user not found")
    }

    if !validatePassword(password, user.Password) {
        return user, errors.New("password does not match")
    }

    return user, nil
}

func hashPassword(password string) string {
    hash := sha256.New()
    hash.Write([]byte(password))
    return hex.EncodeToString(hash.Sum(nil))
}

func validatePassword(providedPassword, storedPassword string) bool {
    return hashPassword(providedPassword) == storedPassword
}

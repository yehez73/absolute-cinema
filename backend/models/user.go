package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
    Admin  Role = "admin"
    Member Role = "member"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string             `bson:"name" json:"name"`
    Email    string             `bson:"email" json:"email"`
    Password string             `bson:"password" json:"password"`
    Phone    string             `bson:"phone" json:"phone"`
    Country  string             `bson:"country" json:"country"`
    Role     Role             `bson:"role" json:"role"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
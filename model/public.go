package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Name      string             `json:"name" bson:"name"`
	Phone     string             `json:"phone" bson:"phone"`
	LineId    string             `json:"lineId" bson:"lineId"`
	IsLogin   bool               `json:"isLogin" bson:"isLogin"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type StateUser []struct {
	Type        string `json:"type" bson:"type"`
	Description string `json:"description" bson:"description"`
	Response    string `json:"response" bson:"response"`
}

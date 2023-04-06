package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	Value       []byte             `json:"value" bson:"value"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type ExpenseDecrypt struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Value       string             `json:"value" bson:"value"`
	Description string             `json:"description" bson:"description"`
}

type Income struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	Value       []byte             `json:"value" bson:"value"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type IncomeDecrypt struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Value       string             `json:"value" bson:"value"`
	Description string             `json:"description" bson:"description"`
}

type RequestTransactions struct {
	Data        string `json:"data" bson:"data"`
	Description string `json:"description" bson:"description"`
}

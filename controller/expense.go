package controller

import (
	"context"
	"example/line-bot-ledger/config"
	"example/line-bot-ledger/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ExpenseCollection *mongo.Collection = config.GetCollection(config.DB, "expenses_transactions")

func GetLastestExpense(userId primitive.ObjectID) (interface{}, error) {
	filter := bson.M{"userId": userId}
	options := options.Find().SetLimit(int64(10))
	cursor, err := ExpenseCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	var results []model.Expense
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetAllExpense(userId primitive.ObjectID) (interface{}, error) {
	filter := bson.M{"userId": userId}
	options := options.Find()
	cursor, err := ExpenseCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	var results []model.Expense
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func CreateExpense(transaction model.Expense) error {
	_, err := ExpenseCollection.InsertOne(context.TODO(), transaction)
	return err
}

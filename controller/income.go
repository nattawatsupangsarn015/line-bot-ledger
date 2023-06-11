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

var IncomeCollection *mongo.Collection = config.GetCollection(config.DB, "incomes_transactions")

func GetLastestIncome(userId primitive.ObjectID) (interface{}, error) {
	limit := int64(10)
	sort := bson.M{"created_at": -1}
	filter := bson.M{"userId": userId}

	options := options.Find().SetSort(sort).SetLimit(limit)
	cursor, err := IncomeCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	var results []model.Income
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetAllIncome(userId primitive.ObjectID) (interface{}, error) {
	filter := bson.M{"userId": userId}

	options := options.Find()
	cursor, err := IncomeCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	var results []model.Income
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func CreateIncome(transaction model.Income) error {
	_, err := IncomeCollection.InsertOne(context.TODO(), transaction)
	return err
}

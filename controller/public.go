package controller

import (
	"context"
	"example/line-bot-ledger/config"
	"example/line-bot-ledger/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = config.GetCollection(config.DB, "users")

func CreateUser(user *model.User) error {
	err := CreateUniqueField(UserCollection, bson.M{"email": 1})
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = UserCollection.InsertOne(context.TODO(), user)
	return err
}

func GetUserByEmail(email string) (interface{}, error) {
	filter := bson.M{"email": email}
	var user model.User
	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func GetUserByLineId(lineId string) (interface{}, error) {
	filter := bson.M{"lineId": lineId, "isLogin": true}
	var user model.User
	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func UpdateLineIdUser(email string, lineId string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"isLogin": true, "lineId": lineId}}
	UserCollection.FindOneAndUpdate(context.TODO(), filter, update)
	return nil
}

package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetUserCollection(collection string) {
	manager.UserCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:UserCollection success!")
}

func (manager *dbManager) FindUserByAccount(account string, opts ...*options.FindOneOptions) models.User {
	filter := bson.M{"account": account}
	var result models.User
	err := manager.UserCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		logrus.Error("manager.FindOneUser error:", err.Error())
		return result
	}
	logrus.Infof("manager.FindOneUser, Found a single document: %+v\n", result)
	return result
}

func (manager *dbManager) FindUserByToken(token string, opts ...*options.FindOneOptions) models.User {
	filter := bson.M{"authorization": token}
	var result models.User
	err := manager.UserCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		logrus.Error("manager.FindOne error:", err.Error())
		return models.User{}
	}
	logrus.Infof("manager.FindOne, Found a single document: %+v\n", result)
	return result
}

func (manager *dbManager) UpdateTokenByAccount(account string, token string, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	filter := bson.M{"account": account}
	update := bson.M{"$set": bson.M{"authorization": token}}
	return manager.UserCollection.UpdateOne(context.Background(), filter, update, opts...)
}

func (manager *dbManager) CreateUser(user models.User) error {
	_, err := manager.UserCollection.InsertOne(context.Background(), user)
	return err
}

func (manager *dbManager) DeleteUserByAccount(account string) (*mongo.DeleteResult, error) {
	filter := bson.M{"account": account}
	return manager.UserCollection.DeleteOne(context.Background(), filter)
}

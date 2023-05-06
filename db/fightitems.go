package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetFightItemsCollection(collection string) {
	manager.FightItemsCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:FightItemsCollection success!")
}

func (manager *dbManager) FindFightItemByFightToken(fightToken int64, opts ...*options.FindOneOptions) models.FightItem {
	filter := bson.M{"fight_token": fightToken}
	var result models.FightItem
	err := manager.FightItemsCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return models.FightItem{}
	}
	return result
}

func (manager *dbManager) FindFightItems(filter bson.M, result *[]models.FightItem, opts ...*options.FindOptions) error {
	cursor, err := manager.FightItemsCollection.Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	// 通过All一次性获取所有结果
	if err = cursor.All(context.Background(), result); err != nil {
		return err
	}

	return nil
}

func (manager *dbManager) CreateFightItem(data models.FightItem, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.FightItemsCollection.InsertOne(context.Background(), data, opts...)
}
package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleSeasonItemsCollection(collection string) {
	manager.RoleSeasonItemsCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleSeasonItemsCollection success!")
}

func (manager *dbManager) FindRoleSeasonItemBySeasonID(roleID int64, seasonID int32, opts ...*options.FindOneOptions) protocols.T_SeasonEntityData {
	filter := bson.M{"role_id": roleID, "season_id": seasonID}
	var result protocols.T_SeasonEntityData
	err := manager.RoleSeasonItemsCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return protocols.T_SeasonEntityData{}
	}
	return result
}

func (manager *dbManager) FindRoleSeasonItemsByRoleID(roleID int64, result *[]protocols.T_SeasonEntityData, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleSeasonItems(filter, result, opts...)
}

func (manager *dbManager) FindRoleSeasonItems(filter bson.M, result *[]protocols.T_SeasonEntityData, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleSeasonItemsCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) InsertOneSeason(data protocols.T_SeasonEntityData, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleSeasonItemsCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleSeasonItemsByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.RoleSeasonItemsCollection.DeleteMany(context.Background(), filter)
}

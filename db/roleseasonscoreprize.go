package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleSeasonScorePrizeCollection(collection string) {
	manager.RoleSeasonScorePrizeCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleSeasonScorePrizeCollection success!")
}

func (manager *dbManager) FindRoleSeasonScorePrizeByRoleID(roleID int64, result *[]models.RoleSeasonScorePrize, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleSeasonScorePrize(filter, result, opts...)
}

func (manager *dbManager) FindRoleSeasonScorePrize(filter bson.M, result *[]models.RoleSeasonScorePrize, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleSeasonForeverScorePrizeCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) InsertOneRoleSeasonScorePrize(data models.RoleSeasonScorePrize, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleSeasonScorePrizeCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleSeasonScorePrizeByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.RoleSeasonScorePrizeCollection.DeleteMany(context.Background(), filter)
}

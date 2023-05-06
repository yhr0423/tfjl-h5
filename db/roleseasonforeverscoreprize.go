package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleSeasonForeverScorePrizeCollection(collection string) {
	manager.RoleSeasonForeverScorePrizeCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleSeasonForverScorePrizeCollection success!")
}

func (manager *dbManager) FindRoleSeasonForeverScorePrizeByRoleID(roleID int64, result *[]models.RoleSeasonForeverScorePrize, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleSeasonForeverScorePrize(filter, result, opts...)
}

func (manager *dbManager) FindRoleSeasonForeverScorePrize(filter bson.M, result *[]models.RoleSeasonForeverScorePrize, opts ...*options.FindOptions) error {
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

func (manager *dbManager) InsertOneRoleSeasonForeverScorePrize(data models.RoleSeasonForeverScorePrize, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleSeasonForeverScorePrizeCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleSeasonForeverScorePrizeByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.RoleSeasonForeverScorePrizeCollection.DeleteMany(context.Background(), filter)
}

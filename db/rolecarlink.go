package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleCarLinkCollection(collection string) {
	manager.RoleCarLinkCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleCarLinkCollection success!")
}

func (manager *dbManager) FindRoleCarLinkByMasterItemID(roleID int64, masterItemID int32, opts ...*options.FindOneOptions) models.RoleCarLink {
	filter := bson.M{"role_id": roleID, "master_item_id": masterItemID}
	var result models.RoleCarLink
	err := manager.RoleCarLinkCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return models.RoleCarLink{}
	}
	return result
}

func (manager *dbManager) FindRoleCarLinkByRoleID(roleID int64, result *[]models.RoleCarLink, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleCarLink(filter, result, opts...)
}

func (manager *dbManager) FindRoleCarLink(filter bson.M, result *[]models.RoleCarLink, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleCarLinkCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) UpdateRoleCarLinkByMasterItemID(roleID int64, masterItemID int32, slaveItemID int32) (*mongo.UpdateResult, error) {
	roleCarLink := manager.FindRoleCarLinkByMasterItemID(roleID, masterItemID)
	if roleCarLink == (models.RoleCarLink{}) {
		roleCarLink = models.RoleCarLink{
			RoleID:       roleID,
			MasterItemID: masterItemID,
			SlaveItemID:  slaveItemID,
		}
		insertOneResult, err := manager.CreateRoleCarLink(roleCarLink)
		if err != nil {
			return nil, err
		}
		return &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 1,
			UpsertedID:    insertOneResult.InsertedID,
		}, nil
	} else {
		filter := bson.M{"role_id": roleID, "master_item_id": masterItemID}
		update := bson.M{"$set": bson.M{"slave_item_id": slaveItemID}}
		return manager.UpdateOneRoleCarLink(filter, update)
	}
}

func (manager *dbManager) UpdateOneRoleCarLink(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return manager.RoleCarLinkCollection.UpdateOne(context.Background(), filter, update, opts...)
}

func (manager *dbManager) CreateRoleCarLink(data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleCarLinkCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleCarLinkByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.DeleteRoleCarLink(filter)
}

func (manager *dbManager) DeleteRoleCarLink(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return manager.RoleCarLinkCollection.DeleteMany(context.Background(), filter, opts...)
}

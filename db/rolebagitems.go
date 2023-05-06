package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleBagItemsCollection(collection string) {
	manager.RoleBagItemsCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleBagItemsCollection success!")
}

func (manager *dbManager) FindCarSkinByItemID(roleID int64, itemID int32, opts ...*options.FindOneOptions) protocols.T_Role_Item {
	filter := bson.M{"role_id": roleID, "id": itemID}
	var result protocols.T_Role_Item
	err := manager.RoleBagItemsCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return protocols.T_Role_Item{}
	}
	return result
}

func (manager *dbManager) FindRoleBagItemsByType(roleID int64, itemType int32, result *[]protocols.T_Role_Item, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID, "type": itemType}
	return manager.FindRoleBagItems(filter, result, opts...)
}

func (manager *dbManager) FindRoleBagItemsByRoleID(roleID int64, result *[]protocols.T_Role_Item, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleBagItems(filter, result, opts...)
}

func (manager *dbManager) FindRoleBagItems(filter bson.M, result *[]protocols.T_Role_Item, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleBagItemsCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) CreateRoleBagItem(data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleBagItemsCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleBagItemByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.RoleBagItemsCollection.DeleteMany(context.Background(), filter)
}

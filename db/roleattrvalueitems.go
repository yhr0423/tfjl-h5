package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleAttrValueItemsCollection(collection string) {
	manager.RoleAttrValueItemsCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleAttrValueItemsCollection success!")
}

func (manager *dbManager) FindRoleAttrValueItemByAttrID(roleID int64, attrID int32, opts ...*options.FindOneOptions) protocols.S_Role_SynRoleAttrValue {
	filter := bson.M{"role_id": roleID, "attr_id": attrID}
	var result protocols.S_Role_SynRoleAttrValue
	err := manager.RoleAttrValueItemsCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return protocols.S_Role_SynRoleAttrValue{}
	}
	return result
}

func (manager *dbManager) FindRoleAttrValueItemsByRoleID(roleID int64, result *[]protocols.S_Role_SynRoleAttrValue, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleAttrValueItems(filter, result, opts...)
}

func (manager *dbManager) FindRoleAttrValueItems(filter bson.M, result *[]protocols.S_Role_SynRoleAttrValue, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleAttrValueItemsCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) UpdateRoleAttrValueByAttrID(roleID int64, attrID int32, value int32) (*mongo.UpdateResult, error) {
	filter := bson.M{"role_id": roleID, "attr_id": attrID}
	update := bson.M{"$set": bson.M{"value": value}}
	return manager.UpdateOneRoleAttrValue(filter, update)
}

func (manager *dbManager) UpdateOneRoleAttrValue(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return manager.RoleAttrValueItemsCollection.UpdateOne(context.Background(), filter, update, opts...)
}

func (manager *dbManager) CreateRoleAttrValueItem(data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleAttrValueItemsCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleAttrValueItemsByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.DeleteRoleAttrValueItems(filter)
}

func (manager *dbManager) DeleteRoleAttrValueItems(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return manager.RoleAttrValueItemsCollection.DeleteMany(context.Background(), filter, opts...)
}

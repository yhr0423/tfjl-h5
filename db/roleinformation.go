package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleInformationCollection(collection string) {
	manager.RoleInformationCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleInformationCollection success!")
}

func (manager *dbManager) FindRoleInformationByRoleID(roleID int64) protocols.T_Information_Data {
	filter := bson.M{"role_id": roleID}
	var result protocols.T_Information_Data
	err := manager.FindOneRoleInformation(filter, &result)
	if err != nil {
		logrus.Error("manager.FindOneRoleInformation error:", err)
		return result
	}
	logrus.Infof("manager.FindOneRoleInformation, Found a single document: %+v\n", result)
	return result
}

func (manager *dbManager) FindOneRoleInformation(filter interface{}, result *protocols.T_Information_Data, opts ...*options.FindOneOptions) error {
	return manager.RoleInformationCollection.FindOne(context.Background(), filter, opts...).Decode(result)
}

func (manager *dbManager) InsertOneRoleInformation(data protocols.T_Information_Data, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleInformationCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleInformationByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.DeleteRoleInfomation(filter)
}

func (manager *dbManager) DeleteRoleInfomation(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return manager.RoleInformationCollection.DeleteOne(context.Background(), filter, opts...)
}

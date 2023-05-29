package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleCollection(collection string) {
	manager.RoleCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleCollection success!")
}

func (manager *dbManager) FindRoleByAccount(account string) models.Role {
	filter := bson.M{"account": account}
	var result models.Role
	err := manager.FindOneRole(filter, &result)
	if err != nil {
		logrus.Error("manager.FindOneRole error:", err)
		return result
	}
	return result
}

func (manager *dbManager) FindRoleByRoleID(roleID int64) models.Role {
	filter := bson.M{"id": roleID}
	var result models.Role
	err := manager.FindOneRole(filter, &result)
	if err != nil {
		logrus.Error("manager.FindOneRole error:", err)
		return result
	}
	return result
}

func (manager *dbManager) FindOneRole(filter interface{}, result *models.Role, opts ...*options.FindOneOptions) error {
	return manager.RoleCollection.FindOne(context.Background(), filter, opts...).Decode(result)
}

func (manager *dbManager) FindRoles(filter interface{}, result *[]models.Role, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) CreateRole(user models.Role) error {
	_, err := manager.RoleCollection.InsertOne(context.Background(), user)
	return err
}

func (manager *dbManager) DeleteRoleByAccount(account string) (int64, error) {
	filter := bson.M{"account": account}
	result, err := manager.RoleCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (manager *dbManager) CountRoles(filter interface{}) (int64, error) {
	return manager.RoleCollection.CountDocuments(context.Background(), filter)
}

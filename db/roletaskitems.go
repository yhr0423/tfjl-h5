package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleTaskItemsCollection(collection string) {
	manager.RoleTaskItemsCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleTaskItemsCollection success!")
}

func (manager *dbManager) FindRoleTaskItemByTaskID(roleID int64, taskID int32, opts ...*options.FindOneOptions) protocols.T_Role_SingleTask {
	filter := bson.M{"role_id": roleID, "task_id": taskID}
	var result protocols.T_Role_SingleTask
	err := manager.RoleTaskItemsCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return protocols.T_Role_SingleTask{}
	}
	return result
}

func (manager *dbManager) FindRoleTaskItemsByRoleID(roleID int64, result *[]protocols.T_Role_SingleTask, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleTaskItems(filter, result, opts...)
}

func (manager *dbManager) FindRoleTaskItems(filter bson.M, result *[]protocols.T_Role_SingleTask, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleTaskItemsCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) InsertOneRoleSingleTask(data protocols.T_Role_SingleTask, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleTaskItemsCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleTaskItemsByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.RoleTaskItemsCollection.DeleteMany(context.Background(), filter)
}

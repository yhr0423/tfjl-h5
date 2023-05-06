package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleBattleArrayCollection(collection string) {
	manager.RoleBattleArrayCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleBattleArrayCollection success!")
}

func (manager *dbManager) FindRoleBattleArrayByRoleID(roleID int64, result *[]protocols.T_Role_BattleArrayIndexData, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleBattleArray(filter, result, opts...)
}

func (manager *dbManager) FindRoleBattleArray(filter bson.M, result *[]protocols.T_Role_BattleArrayIndexData, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleBattleArrayCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) UpdateOneRoleBattleArrayByID(roleID int64, arrayID int32, index int32, heroUUID int64) (*mongo.UpdateResult, error) {
	filter := bson.M{"role_id": roleID, "id": arrayID, "index": index}
	update := bson.M{"$set": bson.M{"hero_uuid": heroUUID}}
	return manager.UpdateOneRoleBattleArray(filter, update)
}

func (manager *dbManager) UpdateOneRoleBattleArray(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return manager.RoleBattleArrayCollection.UpdateOne(context.Background(), filter, update, opts...)
}

func (manager *dbManager) InsertOneRoleBattleArray(data protocols.T_Role_BattleArrayIndexData, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleBattleArrayCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleBattleArrayByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.DeleteRoleBattleArray(filter)
}

func (manager *dbManager) DeleteRoleBattleArray(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return manager.RoleBattleArrayCollection.DeleteMany(context.Background(), filter, opts...)
}

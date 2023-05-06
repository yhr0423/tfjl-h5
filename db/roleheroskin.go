package db

import (
	"context"
	"tfjl-h5/models"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoleHeroSkinCollection(collection string) {
	manager.RoleHeroSkinCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoleHeroSkinCollection success!")
}

func (manager *dbManager) FindRoleHeroSkinByItemUUID(roleID int64, itemUUID int64, opts ...*options.FindOneOptions) models.RoleHeroSkin {
	filter := bson.M{"role_id": roleID, "uuid": itemUUID}
	var result models.RoleHeroSkin
	err := manager.RoleHeroSkinCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return models.RoleHeroSkin{}
	}
	return result
}

func (manager *dbManager) FindRoleHeroSkinByRoleID(roleID int64, result *[]models.RoleHeroSkin, opts ...*options.FindOptions) error {
	filter := bson.M{"role_id": roleID}
	return manager.FindRoleHeroSkin(filter, result, opts...)
}

func (manager *dbManager) FindRoleHeroSkin(filter bson.M, result *[]models.RoleHeroSkin, opts ...*options.FindOptions) error {
	cursor, err := manager.RoleHeroSkinCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) UpdateRoleHeroSkinByItemUUID(roleID int64, itemUUID int64, skinID int32, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	roleHeroSkin := manager.FindRoleHeroSkinByItemUUID(roleID, itemUUID)
	if roleHeroSkin == (models.RoleHeroSkin{}) {
		roleHeroSkin = models.RoleHeroSkin{
			ID_: primitive.NewObjectID(),
			RoleID:     roleID,
			UUID:       itemUUID,
			ID:         skinID,
			CreateTime: int32(time.Now().Unix()),
			Num:        1,
		}
		insertOneResult, err := manager.InsertOneRoleHeroSkin(roleHeroSkin)
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
		filter := bson.M{"role_id": roleID, "uuid": itemUUID}
		update := bson.M{"$set": bson.M{"id": skinID}}
		return manager.RoleHeroSkinCollection.UpdateOne(context.Background(), filter, update, opts...)
	}
}

func (manager *dbManager) InsertOneRoleHeroSkin(data models.RoleHeroSkin, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.RoleHeroSkinCollection.InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) DeleteRoleHeroSkinByRoleID(roleID int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"role_id": roleID}
	return manager.RoleHeroSkinCollection.DeleteMany(context.Background(), filter)
}
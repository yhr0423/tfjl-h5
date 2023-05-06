package db

import (
	"context"
	"tfjl-h5/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetRoomCollection(collection string) {
	manager.RoomCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:RoomCollection success!")
}

func (manager *dbManager) FindRoomByRoomIDFightType(roomID string, fightType int32, opts ...*options.FindOneOptions) models.Room {
	filter := bson.M{"room_id": roomID, "fight_type": fightType, "status": 0}
	var result models.Room
	err := manager.RoomCollection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		return models.Room{}
	}
	return result
}

func (manager *dbManager) FindRoomsByStatus(roleID int64, status int32, result *[]models.Room, opts ...*options.FindOptions) error {
	filter := bson.M{"creator_role_id": roleID, "status": status}
	return manager.FindRooms(filter, result, opts...)
}

func (manager *dbManager) FindRoomsByRoleID(roleID int64, result *[]models.Room, opts ...*options.FindOptions) error {
	filter := bson.M{"creator_role_id": roleID}
	return manager.FindRooms(filter, result, opts...)
}

func (manager *dbManager) FindRooms(filter bson.M, result *[]models.Room, opts ...*options.FindOptions) error {
	cursor, err := manager.RoomCollection.Find(context.Background(), filter, opts...)
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

func (manager *dbManager) UpdateRoomStatus(ObjectID interface{}, status int32) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": bson.M{"status": status}}
	return manager.UpdateOneRoom(filter, update)
}

func (manager *dbManager) UpdateOneRoom(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return manager.RoomCollection.UpdateOne(context.Background(), filter, update, opts...)
}

func (p *dbManager) CreateRoom(roomID string, longRoomID string, creatorRoleID int64, fightType int32, status int32) error {
	data := models.Room{
		ID_:           primitive.NewObjectID(),
		RoomID:        roomID,
		LongRoomID:    longRoomID,
		CreatorRoleID: creatorRoleID,
		FightType:     fightType,
		Status:        status,
	}
	_, err := p.InsertOneRoom(data)
	return err
}

func (p *dbManager) InsertOneRoom(data models.Room, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return p.RoomCollection.InsertOne(context.Background(), data, opts...)
}

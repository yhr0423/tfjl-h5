package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID_           primitive.ObjectID `bson:"_id"`
	RoomID        string             `bson:"room_id"`
	LongRoomID    string             `bson:"long_room_id"`
	CreatorRoleID int64              `bson:"creator_role_id"`
	FightType     int32              `bson:"fight_type"`
	Status        int32              `bson:"status"` // 0:未开始 1:进行中 2:已结束 3:已销毁
}

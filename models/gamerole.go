package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID_                      primitive.ObjectID `bson:"_id"`                         // 数据库ID
	RoleID                   int64              `bson:"id"`                          // 角色ID
	Account                  string             `bson:"account"`                     // 角色绑定的账号
	StrID                    string             `bson:"show_id"`                     // 显示ID
	RoleName                 string             `bson:"name"`                        // 角色名字
	BIndulge                 bool               `bson:"b_indulge"`                   // 是否处于防沉迷状态
	IndulgeTime              int32              `bson:"indulge_time"`                // 防沉迷时间
	IndulgeDayOnlineTime     int32              `bson:"indulge_day_online_time"`     // 防沉迷当天在线时间
	Key                      uint8              `bson:"key"`                         // 数据加密key
	Level                    int32              `bson:"level"`                       // 等级
	ForbidLoginTimeRemaining int32              `bson:"forbid_login_time_remaining"` // 封禁剩余时间
	BattleArraySelectID      int32              `bson:"battle_array_select_id"`      // 封禁剩余时间
}

type RoleBattleArray struct {
	RoleID     int64  `bson:"role_id"`
	ID         int32  `bson:"id"`
	Index      int32  `bson:"index"`
	HeroUUID   int64  `bson:"hero_uuid"`
	BattleName string `bson:"battle_name"`
}

type RoleHeroSkin struct {
	RoleID     int64  `bson:"role_id"`
	UUID       int64  `bson:"uuid"`
	ID         int32  `bson:"id"`
	CreateTime int32  `bson:"create_time"`
	Name       string `bson:"name"`
	Num        int32  `bson:"num"`
}

type RoleSeasonForeverScorePrize struct {
	ID_     primitive.ObjectID `bson:"_id"`
	RoleID  int64              `bson:"role_id"`
	ID      int32              `bson:"id"`
	IsPrize bool               `bson:"is_prize"`
	IsExtra bool               `bson:"is_extra"`
}

type RoleSeasonScorePrize struct {
	ID_     primitive.ObjectID `bson:"_id"`
	RoleID  int64              `bson:"role_id"`
	ID      int32              `bson:"id"`
	IsPrize bool               `bson:"is_prize"`
	IsExtra bool               `bson:"is_extra"`
}

type RoleCarLink struct {
	RoleID       int64 `bson:"role_id"`
	MasterItemID int32 `bson:"master_item_id"`
	SlaveItemID  int32 `bson:"slave_item_id"`
}

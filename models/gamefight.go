package models

type FightItem struct {
	FightToken  string  `bson:"fight_token"`
	Roles       []int64 `bson:"roles"`
	FightStatus int32   `bson:"fight_status"`
}
